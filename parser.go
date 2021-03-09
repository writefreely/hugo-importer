package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/writeas/web-core/i18n"
)

var wd string = ""

func ParseContentDirectory(p string, s bool) ([]PostToMigrate, error) {
	var numberOfFiles int = 0

	// Get the current working directory.
	rwd, err := os.Getwd()
	wd = rwd

	// Find the config file
	matches, err := filepath.Glob("config.*")
	if err != nil {
		log.Fatal(err)
	}
	fileComponents := strings.Split(matches[0], ".")
	format := fileComponents[len(fileComponents)-1]
	configFilePath := filepath.Join(rwd, matches[0])
	languageCode, err := ScanConfigForLanguage(configFilePath, format)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Setting language:", languageCode)
	baseURL, err := ScanConfigForBaseUrl(configFilePath, format)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found baseURL:", baseURL)

	// Change directory to the path passed in.
	srcPath := filepath.Join(rwd, "content", p)
	os.Chdir(srcPath)
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	var posts []PostToMigrate
	filepath.Walk(wd, func(path string, i os.FileInfo, err error) error {
		if !i.IsDir() && !strings.HasPrefix(i.Name(), ".") && strings.HasSuffix(i.Name(), ".md") {
			fileName := i.Name()
			fmt.Println("> Parsing", fileName)
			post, err := parsePost(fileName, languageCode, baseURL, s)
			if err != nil {
				log.Fatal(err)
			}
			posts = append(posts, post)
			fmt.Println("> Finished parsing", fileName)
			numberOfFiles += 1
		}
		return nil
	})

	fmt.Printf("Parsed %d files\n\n", numberOfFiles)

	// Change the working directory back to /content
	os.Chdir(rwd)
	rwd, err = os.Getwd()

	return posts, nil
}

func parsePost(f string, l string, b string, s bool) (PostToMigrate, error) {
	file, err := os.Open(f)

	if err != nil {
		return PostToMigrate{}, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	pf, err := pageparser.ParseFrontMatterAndContent(file)

	hashtags := []string{}

	var post PostToMigrate
	var created time.Time
	var updated time.Time

	if pf.FrontMatterFormat == metadecoders.JSON ||
		pf.FrontMatterFormat == metadecoders.YAML ||
		pf.FrontMatterFormat == metadecoders.TOML {
		for k, v := range pf.FrontMatter {
			if k == "date" {
				c, err := time.Parse(time.RFC3339, v.(string))
				if err != nil {
					return PostToMigrate{}, err
				}
				created = c.UTC()
			}
			if k == "lastMod" {
				u, err := time.Parse(time.RFC3339, v.(string))
				if err != nil {
					return PostToMigrate{}, err
				}
				updated = u.UTC()
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
						hashtags = append(hashtags, ConvertToHashtag(t))
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
						hashtags = append(hashtags, ConvertToHashtag(c))
					}
				}
			}
		}
		content := string(pf.Content[:]) + "\n\n" + strings.Join(hashtags, " ")
		if s {
			content = scanContentForLocalImages(content, b)
		}

		var slug string
		if pf.FrontMatter["slug"] != nil {
			slug = pf.FrontMatter["slug"].(string)
		} else {
			slug = ""
		}

		rtl := i18n.LangIsRTL(l)

		post = PostToMigrate{
			body:    content,
			title:   pf.FrontMatter["title"].(string),
			slug:    slug,
			lang:    &l,
			rtl:     &rtl,
			created: &created,
			updated: &updated,
		}
	}

	return post, nil
}

func scanContentForLocalImages(c string, b string) string {
	// Search for Markdown image links with optional alt text
	var reMarkdown = regexp.MustCompile(`!\[.*\((?P<url>.+)\)`)
	mdMatches := reMarkdown.FindAllStringSubmatch(c, -1)
	for _, mdMatch := range mdMatches {
		img := mdMatch[1]
		if ImageIsLocal(img, b) {
			// Strip the base URL if the post uses an absolute URL.
			if strings.HasPrefix(img, b) {
				img = strings.Replace(img, b, "", 1)
			}
			imgURL := uploadOrLogError(img)
			c = strings.Replace(c, img, imgURL, -1)
		} else {
			fmt.Println("  > 🖼 (⛔️) Skipping upload of remote image: ", img)
		}
	}

	// Search for HTML image links with optional alt text
	var reHtml = regexp.MustCompile(`<img.*src="(?P<url>\S+)".*/>`)
	htmlMatches := reHtml.FindAllStringSubmatch(c, -1)
	for _, htmlMatch := range htmlMatches {
		img := htmlMatch[1]
		if ImageIsLocal(img, b) {
			// Strip the base URL if the post uses an absolute URL.
			if strings.HasPrefix(img, b) {
				img = strings.Replace(img, b, "", 1)
			}
			imgURL := uploadOrLogError(img)
			c = strings.Replace(c, img, imgURL, -1)
		} else {
			fmt.Println("  > 🖼 (⛔️) Skipping upload of remote image: ", img)
		}
	}

	return c
}

func uploadOrLogError(i string) string {
	ip := filepath.Join(wd, "static", i)
	fmt.Println("  > 🖼 (⏳) Uploading image to Snap.as:", ip)

	retries := 3
	var upErr string

	for retries > 0 {
		imgURL, err := UploadImage(ip)
		if err != nil {
			upErr = err.Error()
			retries--
			fmt.Println("  > 🖼 (⚠️) Upload failed. Retrying...")
		} else {
			fmt.Println("  > 🖼 (✅) Upload complete, at:", imgURL)
			return imgURL
		}
	}
	fmt.Println("  > 🖼 (⚠️) Upload failed. Logging error and skipping.")
	LogUploadError(i, upErr)
	return i
}

type PostToMigrate struct {
	body    string
	title   string
	slug    string
	lang    *string
	rtl     *bool
	created *time.Time
	updated *time.Time
}
