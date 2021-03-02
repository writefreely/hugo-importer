package main

import (
	"github.com/writeas/go-writeas/v2"
)

func PublishPost(p PostToMigrate, a string, c *writeas.Client) error {
	wa, err := c.CreatePost(&writeas.PostParams{
		Title:      p.title,
		Content:    p.body,
		Collection: a,
		Slug:       p.slug,
		Created:    p.created,
		Updated:    p.updated,
		IsRTL:      p.rtl,
		Language:   p.lang,
	})

	if err != nil {
		return err
	}

	LogResponse(wa)

	return nil
}
