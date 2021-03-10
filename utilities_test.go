package main

import (
	"strings"
	"testing"
)

func TestSplitAnyStringWithMultipleSeparators(t *testing.T) {
	// Given
	testString := "hello_there wild-world.txt"
	separators := " -_."
	wanted := "hellotherewildworldtxt"

	// When
	result := strings.Join(SplitAny(testString, separators), "")

	// Then
	if result != "hellotherewildworldtxt" {
		t.Fatalf(`SplitAny() = %s, want %s`, result, wanted)
	}
}

func TestConvertToHashtagOnSingleWordString(t *testing.T) {
	// Given
	testString := "potatoes"
	wanted := "#potatoes"

	// When
	result := ConvertToHashtag(testString)

	// Then
	if result != wanted {
		t.Fatalf(`ConvertToHashtag() = %s, want %s`, result, wanted)
	}
}

func TestConvertToHashtagOnMultipleWordString(t *testing.T) {
	// Given
	testString := "baked-potatoes"
	wanted := "#bakedPotatoes"

	// When
	result := ConvertToHashtag(testString)

	// Then
	if result != wanted {
		t.Fatalf(`ConvertToHashtag() = %s, want %s`, result, wanted)
	}
}

func TestImageIsLocalIfRelativePath(t *testing.T) {
	// Given
	testPath := "/images/picture.jpg"
	baseURL := "https://example.com"
	wanted := true

	// When
	result := ImageIsLocal(testPath, baseURL)

	// Then
	if result != wanted {
		t.Fatalf(`ImageIsLocal() = %t, want %t`, result, wanted)
	}
}

func TestImageIsLocalIfAbsolutePathWithSameBaseURL(t *testing.T) {
	// Given
	testPath := "https://example.com/images/picture.jpg"
	baseURL := "https://example.com"
	wanted := true

	// When
	result := ImageIsLocal(testPath, baseURL)

	// Then
	if result != wanted {
		t.Fatalf(`ImageIsLocal() = %t, want %t`, result, wanted)
	}
}

func TestImageIsLocalIfAbsolutePathWithDifferentBaseURL(t *testing.T) {
	// Given
	testPath := "https://example.com/images/picture.jpg"
	baseURL := "https://write.as"
	wanted := false

	// When
	result := ImageIsLocal(testPath, baseURL)

	// Then
	if result != wanted {
		t.Fatalf(`ImageIsLocal() = %t, want %t`, result, wanted)
	}
}

func TestScanConfigForBaseUrlWithValidConfigFile(t *testing.T) {
	// Given
	configFilePath := "testdata/config-valid.toml"
	configFileFormat := "toml"
	wanted := "https://example.com/"

	// When
	result, err := ScanConfigForBaseUrl(configFilePath, configFileFormat)
	if err != nil {
		t.Fatalf(`ScanConfigForBaseUrl() threw error %s`, err)
	}

	// Then
	if result != wanted {
		t.Fatalf(`ScanConfigForBaseUrl() = %s, want %s`, result, wanted)
	}
}

func TestScanConfigForBaseUrlWithInvalidConfigFile(t *testing.T) {
	// Given
	configFilePath := "testdata/config-noBaseUrl.toml"
	configFileFormat := "toml"

	// When
	result, err := ScanConfigForBaseUrl(configFilePath, configFileFormat)

	// Then
	if err == nil {
		t.Fatalf(`ScanConfigForBaseUrl() = %s, want error`, result)
	}
}

func TestScanConfigForLanguageWithValidConfigFile(t *testing.T) {
	// Given
	configFilePath := "testdata/config-valid.toml"
	configFileFormat := "toml"
	wanted := "en"

	// When
	result, err := ScanConfigForLanguage(configFilePath, configFileFormat)
	if err != nil {
		t.Fatalf(`ScanConfigForBaseUrl() threw error %s`, err)
	}

	// Then
	if result != wanted {
		t.Fatalf(`ScanConfigForLanguage() = %s, want %s`, result, wanted)
	}
}

func TestScanConfigForLanguageWithInvalidConfigFile(t *testing.T) {
	// Given
	configFilePath := "testdata/config-noLang.toml"
	configFileFormat := "toml"

	// When
	result, err := ScanConfigForLanguage(configFilePath, configFileFormat)

	// Then
	if err == nil {
		t.Fatalf(`ScanConfigForLanguage() = %s, want error`, result)
	}
}
