package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/gohugoio/hugo/parser/pageparser"
)

func ParseContentDirectory(p string) error {
	var numberOfFiles int = 0

	// Get the current working directory.
	rwd, err := os.Getwd()

	// Change directory to the path passed in.
	os.Chdir(rwd + "/" + p);
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(wd, func(path string, i os.FileInfo, err error) error {
		if !i.IsDir() && !strings.HasPrefix(i.Name(), ".") && strings.HasSuffix(i.Name(), ".md") {
			fileName := i.Name()
			fmt.Println("> Parsing", fileName)
			parsePost(fileName)
			fmt.Println("> Finished parsing", fileName)
			numberOfFiles += 1
		}
		return nil
	})

	fmt.Printf("Parsed %d files\n\n", numberOfFiles)

	// Change the working directory back to /content
	os.Chdir(rwd);
	rwd, err = os.Getwd()

	return nil
}

func parsePost(f string) error {
	file, err := os.Open(f)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	pf, err := pageparser.ParseFrontMatterAndContent(file)

	hashtags := []string{}

	if 
		pf.FrontMatterFormat == metadecoders.JSON ||
		pf.FrontMatterFormat == metadecoders.YAML ||
		pf.FrontMatterFormat == metadecoders.TOML {
		for k, v := range pf.FrontMatter {
			switch vv := v.(type) {
			case time.Time:
				pf.FrontMatter[k] = vv.Format(time.RFC3339)
			}
			if k == "tags" {
				// Enumerate the tags, translate them to hashtags,
				// and append them to hastags array
				tags, ok := v.([]interface{})
				if !ok {
					continue
				}
				for _, ti := range tags {
					if t, ok := ti.(string); ok {
						hashtags = append(hashtags, convertToHashtag(t))
					}
				}
			} else if k == "categories" {
				// Enumerate the categories, translate them to hashtags,
				// and append them to hashtag array
				categories, ok := v.([]interface{})
				if !ok {
					continue
				}
				for _, ci := range categories {
					if c, ok := ci.(string); ok {
						hashtags = append(hashtags, convertToHashtag(c))
					}
				}
			}
		}
		content := string(pf.Content[:]) + "\n\n" + strings.Join(hashtags, " ")

		post := postToMigrate{
			title: pf.FrontMatter["title"].(string),
			created: pf.FrontMatter["date"].(string),
			body: content,
		}
		fmt.Printf("> Title: %+v\n", post.title)
		fmt.Printf("> Published: %+v\n", post.created)
	}

	return nil
}

func convertToHashtag(s string) string {
	hashtagPrefix := "#"
	words := SplitAny(s, " -_.")

	// Collapse the words array to a single, camelCased string,
	// and prefix with an octothorpe
	if len(words) > 1 {
		for i := 1; i < len(words); i++ {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}
		return hashtagPrefix + strings.Join(words, "")
	} else {
		return hashtagPrefix + words[0]
	}
}

type postToMigrate struct {
	title string
	created string
	body string
}
