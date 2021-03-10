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