package main

import (
	"fmt"
	"log"

	"github.com/writeas/go-writeas/v2"
)

func SignIn(u string, p string) (*writeas.Client, error) {
	fmt.Println("Logging in...")
	c := writeas.NewClient()
	_, err := c.LogIn(u, p)
	if err != nil {
		return nil, err
	}
	fmt.Println("Logged in!")
	return c, nil
}

func SignOut(c *writeas.Client) {
	fmt.Println("Logging out...")
	err := c.LogOut()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logged out!")
}
