package main

import (
	"testing"
	"time"
)

func TestScanContentForShortCodesForSimpleGist(t *testing.T) {
	// Given
	shortcode := "{{< gist spf13 7896402 >}}"
	wanted := "https://gist.github.com/spf13/7896402"

	// When
	result := scanContentForShortcodes(shortcode)

	// Then
	if result != wanted {
		t.Fatalf(`scanContentForShortcodes() = %s, want %s`, result, wanted)
	}
}

func TestScanContentForShortCodesForComplexGist(t *testing.T) {
	// Given
	shortcode := "{{< gist spf13 7896402 \"img.html\" >}}"
	wanted := "https://gist.github.com/spf13/7896402"

	// When
	result := scanContentForShortcodes(shortcode)

	// Then
	if result != wanted {
		t.Fatalf(`scanContentForShortcodes() = %s, want %s`, result, wanted)
	}
}

func TestScanContentForShortCodesForSimpleInstagram(t *testing.T) {
	// Given
	shortcode := "{{< instagram BWNjjyYFxVx >}}"
	wanted := "https://www.instagram.com/p/BWNjjyYFxVx"

	// When
	result := scanContentForShortcodes(shortcode)

	// Then
	if result != wanted {
		t.Fatalf(`scanContentForShortcodes() = %s, want %s`, result, wanted)
	}
}

func TestScanContentForShortCodesForComplexInstagram(t *testing.T) {
	// Given
	shortcode := "{{< instagram BWNjjyYFxVx hidecaption >}}"
	wanted := "https://www.instagram.com/p/BWNjjyYFxVx"

	// When
	result := scanContentForShortcodes(shortcode)

	// Then
	if result != wanted {
		t.Fatalf(`scanContentForShortcodes() = %s, want %s`, result, wanted)
	}
}

func TestScanContentForShortCodesForTwitter(t *testing.T) {
	// Given
	shortcode := "{{< tweet 877500564405444608 >}}"
	wanted := "https://twitter.com/twitter/status/877500564405444608"

	// When
	result := scanContentForShortcodes(shortcode)

	// Then
	if result != wanted {
		t.Fatalf(`scanContentForShortcodes() = %s, want %s`, result, wanted)
	}
}

func TestScanContentForShortCodesForVimeo(t *testing.T) {
	// Given
	shortcode := "{{< vimeo 146022717 >}}"
	wanted := "https://player.vimeo.com/video/146022717"

	// When
	result := scanContentForShortcodes(shortcode)

	// Then
	if result != wanted {
		t.Fatalf(`scanContentForShortcodes() = %s, want %s`, result, wanted)
	}
}

func TestScanContentForShortCodesForSimpleYoutube(t *testing.T) {
	// Given
	shortcode := "{{< youtube w7Ft2ymGmfc >}}"
	wanted := "https://www.youtube.com/watch?v=w7Ft2ymGmfc"

	// When
	result := scanContentForShortcodes(shortcode)

	// Then
	if result != wanted {
		t.Fatalf(`scanContentForShortcodes() = %s, want %s`, result, wanted)
	}
}

func TestScanContentForShortCodesForComplexYoutube(t *testing.T) {
	// Given
	shortcode := "{{< youtube id=\"w7Ft2ymGmfc\" autoplay=\"true\" >}}"
	wanted := "https://www.youtube.com/watch?v=w7Ft2ymGmfc"

	// When
	result := scanContentForShortcodes(shortcode)

	// Then
	if result != wanted {
		t.Fatalf(`scanContentForShortcodes() = %s, want %s`, result, wanted)
	}
}

func TestParsePostWithValidContent(t *testing.T) {
	// Given
	postFilePath := "testdata/sample-post.md"
	language := "en"
	baseUrl := "https://example.com"
	scanForImages := false
	rtl := false
	created, _ := time.Parse(time.RFC3339, "2021-03-01T09:00:00-04:00")
	var updated time.Time
	title := "Sample Post"
	content := "\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc vitae tellus tempus, convallis sapien vitae, ullamcorper odio.\n\nhttps://gist.github.com/spf13/7896402\n\nAliquam vulputate cursus dictum. Integer semper ipsum et ligula euismod, vel dignissim ante consectetur. Donec vel auctor nisi.\n\n#firstCategory #categoryTwo #options #moreOptions"
	
	wanted := PostToMigrate{
		body: content,
		title: title,
		lang: &language,
		rtl: &rtl,
		created: &created,
		updated: &updated,
	}

	// When
	result, _ := parsePost(postFilePath, language, baseUrl, scanForImages)

	// Then
	if result == wanted {
		t.Fatalf(`parsePost() failed to parse content correctly`)
	}
}
