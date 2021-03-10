package main

import (
	"github.com/snapas/go-snapas"
	"github.com/writeas/go-writeas/v2"
)

func PublishPost(p PostToMigrate, a string) error {
	wa, err := Client.CreatePost(&writeas.PostParams{
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

func UploadImage(p string) (string, error) {
	t := Client.Token()
	// sc := snapas.NewDevClient(t)		// Use in development
	sc := snapas.NewClient(t)
	i, err := sc.UploadPhoto(&snapas.PhotoParams{
		FileName: p,
	})
	if err != nil {
		return "", err
	}
	return i.URL, nil
}
