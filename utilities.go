package main

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

// Credit: https://stackoverflow.com/a/54426140
func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func ConvertToHashtag(s string) string {
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

func ImageIsLocal(p string, b string) bool {
	// If the path starts with http, check to see if it's local or remote
	if strings.HasPrefix(p, "http") {
		return strings.HasPrefix(p, b)
	}
	// If it doesn't start with http, it's local, so we return true.
	return true
}

func ScanConfigForBaseUrl(p string, f string) (string, error) {
	var baseURL string

	var format metadecoders.Format

	switch f {
	case "json":
		format = metadecoders.JSON
	case "toml":
		format = metadecoders.TOML
	case "yaml":
		format = metadecoders.YAML
	default:
		err := errors.New("Invalid config file format found")
		return "", err
	}

	content, err := ioutil.ReadFile(p)
	if err != nil {
		return "", err
	}

	m, err := metadecoders.Default.UnmarshalToMap(content, format)
	if err != nil {
		return "", err
	}

	if m["baseURL"] != nil && len(m["baseURL"].(string)) > 0 {
		baseURL = m["baseURL"].(string)
	} else {
		return "", errors.New("No baseURL value found")
	}

	return baseURL, nil
}

func ScanConfigForLanguage(p string, f string) (string, error) {
	var languageCode string

	var format metadecoders.Format

	switch f {
	case "json":
		format = metadecoders.JSON
	case "toml":
		format = metadecoders.TOML
	case "yaml":
		format = metadecoders.YAML
	default:
		err := errors.New("Invalid config file format found")
		return "", err
	}

	content, err := ioutil.ReadFile(p)
	if err != nil {
		return "", err
	}

	m, err := metadecoders.Default.UnmarshalToMap(content, format)
	if err != nil {
		return "", err
	}

	if m["languageCode"] != nil {
		languageCode = m["languageCode"].(string)
	} else {
		return "", errors.New("No language found")
	}
	if m["defaultContentLanguage"] != nil {
		languageCode = m["defaultContentLanguage"].(string)
	}

	if len(languageCode) == 0 {
		return "", errors.New("No language found")
	}

	return languageCode[0:2], nil
}
