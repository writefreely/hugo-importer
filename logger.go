package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/writeas/go-writeas/v2"
)

var responses = []response{}

func LogResponse(p *writeas.Post) {
	r := response{
		ID:         p.ID,
		Slug:       p.Slug,
		Font:       p.Font,
		Language:   p.Language,
		RTL:        p.RTL,
		Created:    p.Created,
		Updated:    p.Updated,
		Title:      p.Title,
		Content:    p.Content,
		Tags:       p.Tags,
		Collection: p.Collection.Alias,
	}
	responses = append(responses, r)
}

func WriteResponsesToDisk() error {
	if len(responses) == 0 {
		return errors.New("Could not find publishing results")
	}

	responsesJson, _ := json.Marshal(responses)
	fmt.Println("Writing publishing log to disk...")

	t := time.Now()
	ts := t.Format("20060102150405")
	f := "publishing-log_" + ts + ".json"

	err := ioutil.WriteFile(f, responsesJson, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Publishing log written.")
	return nil
}

type response struct {
	ID         string    `json:"id"`
	Slug       string    `json:"slug"`
	Font       string    `json:"appearance"`
	Language   *string   `json:"language"`
	RTL        *bool     `json:"rtl"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	Title      string    `json:"title"`
	Content    string    `json:"body"`
	Tags       []string  `json:"tags"`
	Collection string    `json:"collection,omitempty"`
}
