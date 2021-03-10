package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/writeas/go-writeas/v2"
)

var responses = []response{}
var uploadErrors = []uploadError{}

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

func LogUploadError(i string, e string) {
	ue := uploadError{
		ImagePath:   i,
		UploadError: e,
	}
	uploadErrors = append(uploadErrors, ue)
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

func WriteUploadErrorsTo(p string) error {
	if len(uploadErrors) > 0 {
		fmt.Println("Writing image-upload error log to disk...")
		f := "upload-error.log"
		logfilePathName := filepath.Join(p, f)
		file, err := os.Create(logfilePathName)
		if err != nil {
			return err
		}
		defer file.Close()

		w := bufio.NewWriter(file)
		for _, line := range uploadErrors {
			fmt.Fprintln(w, line)
		}
		fmt.Println("Error log written.")
		return w.Flush()
	}
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

type uploadError struct {
	ImagePath   string
	UploadError string
}
