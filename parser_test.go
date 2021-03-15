package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScanContentForShortCodes(t *testing.T) {
	testCases := map[string]struct {
		given    string
		expected string
	}{
		"simple gist":           {"{{< gist spf13 7896402 >}}", "https://gist.github.com/spf13/7896402"},
		"complex gist":          {"{{< gist spf13 7896402 \"img.html\" >}}", "https://gist.github.com/spf13/7896402"},
		"simple instagram":      {"{{< instagram BWNjjyYFxVx >}}", "https://www.instagram.com/p/BWNjjyYFxVx"},
		"complex instagram":     {"{{< instagram BWNjjyYFxVx hidecaption >}}", "https://www.instagram.com/p/BWNjjyYFxVx"},
		"twitter":               {"{{< tweet 877500564405444608 >}}", "https://twitter.com/twitter/status/877500564405444608"},
		"vimeo":                 {"{{< tweet 877500564405444608 >}}", "https://twitter.com/twitter/status/877500564405444608"},
		"simple youtube":        {"{{< youtube w7Ft2ymGmfc >}}", "https://www.youtube.com/watch?v=w7Ft2ymGmfc"},
		"complex youtube":       {"{{< youtube id=\"w7Ft2ymGmfc\" autoplay=\"true\" >}}", "https://www.youtube.com/watch?v=w7Ft2ymGmfc"},
		"unsupported shortcode": {"{{< highlight html >}}", "{{< highlight html >}}"},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result := scanContentForShortcodes(test.given)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestParsePostWithValidContent(t *testing.T) {
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
		body:    content,
		title:   title,
		lang:    &language,
		rtl:     &rtl,
		created: &created,
		updated: &updated,
	}

	result, _ := parsePost(postFilePath, language, baseUrl, scanForImages)

	if result == wanted {
		t.Fatalf(`parsePost() failed to parse content correctly`)
	}
}
