package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/writeas/go-writeas/v2"
)

var Client *writeas.Client

func SignIn(u string, p string, i string) error {
	if i == "" {
		fmt.Println("Logging in...")
		Client = writeas.NewClient()
		_, err := Client.LogIn(u, p)
		if err != nil {
			return err
		}
		fmt.Println("Logged in!")
		return nil
	}

	instance, err := url.Parse(i)
	if err != nil {
		return err
	}
	instance.Scheme = "https"
	instance.Path += "/api"

	fmt.Println("Logging in to", i)
	config := writeas.Config{URL: instance.String()}
	Client = writeas.NewClientWith(config)
	_, err = Client.LogIn(u, p)
	if err != nil {
		return err
	}
	fmt.Println("Logged in!")
	return nil
}

func SignOut() {
	fmt.Println("Logging out...")
	err := Client.LogOut()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logged out!")
}
