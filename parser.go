package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	// "github.com/gohugoio/hugo/parser"
	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/urfave/cli/v2"
)



var cmdParseSource = cli.Command{
	Name: "parse",
	Usage: "Parse the content source directory",
	Action: parseContentDirectory,
}

func parseContentDirectory(c *cli.Context) error {
	var numberOfFiles int = 0

	wd, err := os.Getwd()
	fmt.Println("Working Directory ->", wd)

	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(wd, func(path string, i os.FileInfo, err error) error {
		if !i.IsDir() && !strings.HasPrefix(i.Name(), ".") && strings.HasSuffix(i.Name(), ".md") {
			fileName := i.Name()
			fmt.Println("> Parsing", fileName)
			parsePost(fileName)
			fmt.Println("> Finished parsing", fileName)
			fmt.Println("")
			numberOfFiles += 1
		}
		return nil
	})

	fmt.Printf("Parsed %d files", numberOfFiles)

	return nil
}

func parsePost(f string) error {
	// hashtags := []string{}

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

	if pf.FrontMatterFormat == metadecoders.JSON || pf.FrontMatterFormat == metadecoders.YAML || pf.FrontMatterFormat == metadecoders.TOML {
		fmt.Println("> Parsing front matter...")
		for k, v := range pf.FrontMatter {
			switch vv := v.(type) {
			case time.Time:
				pf.FrontMatter[k] = vv.Format(time.RFC3339)
			}
			if k == "tags" {
				// Enumerate the tags, translate them to hashtags, and append them to hastags array
			}
			if k == "categories" {
				// Enumerate the categories, translate them to hashtags, and append them to hashtag array
			}
			fmt.Println(k, "->", v)
		}
		fmt.Println("> Front matter parsed")
	}

	return nil
}