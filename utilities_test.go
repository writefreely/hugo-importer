package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitAny(t *testing.T) {
	testCases := map[string]struct {
		givenString     string
		givenSeparators string
		expected        string
	}{
		"multiple_separators": {"hello_there wild-world.txt", " -_.", "hellotherewildworldtxt"},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result := strings.Join(SplitAny(test.givenString, test.givenSeparators), "")
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestConvertToHashtag(t *testing.T) {
	testCases := map[string]struct {
		given    string
		expected string
	}{
		"single-word string":   {"potatoes", "#potatoes"},
		"multiple-word string": {"baked-potatoes", "#bakedPotatoes"},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result := ConvertToHashtag(test.given)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestImageIsLocal(t *testing.T) {
	testCases := map[string]struct {
		givenTestPath string
		givenBaseURL  string
		expected      bool
	}{
		"relative path":                       {"/images/picture.jpg", "https://example.com", true},
		"absolute path w/ same base URL":      {"https://example.com/images/picture.jpg", "https://example.com", true},
		"absolute path w/ different base URL": {"https://example.com/images/picture.jpg", "https://write.as", false},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result := ImageIsLocal(test.givenTestPath, test.givenBaseURL)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestScanConfigForBaseUrl(t *testing.T) {
	testCases := map[string]struct {
		givenConfigPath   string
		givenConfigFormat string
		expected          string
	}{
		"valid config file":   {"testdata/config-valid.toml", "toml", "https://example.com/"},
		"invalid config file": {"testdata/config-noBaseUrl.toml", "toml", "No baseURL value found"},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := ScanConfigForBaseUrl(test.givenConfigPath, test.givenConfigFormat)
			if err != nil {
				assert.EqualError(t, err, test.expected)
			} else {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestScanConfigForLanguage(t *testing.T) {
	testCases := map[string]struct {
		givenConfigPath   string
		givenConfigFormat string
		expected          string
	}{
		"valid config file":   {"testdata/config-valid.toml", "toml", "en"},
		"invalid config file": {"testdata/config-noLang.toml", "toml", "No language found"},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := ScanConfigForLanguage(test.givenConfigPath, test.givenConfigFormat)
			if err != nil {
				assert.EqualError(t, err, test.expected)
			} else {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
