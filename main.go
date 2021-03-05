package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

func main() {
	var username string
	var dstBlog string
	var srcPath string
	var uploadImages bool
	var instanceUrl string

	app := &cli.App{
		Name:  "Write.as Hugo Importer",
		Usage: "Import a Hugo source directory into Write.as/WriteFreely by running this importer from the root directory of your Hugo site.",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "user",
				Aliases:     []string{"u"},
				Usage:       "The username for the Write.as/WriteFreely account",
				Required:    true,
				Destination: &username,
			},

			&cli.StringFlag{
				Name:        "blog",
				Aliases:     []string{"b"},
				Usage:       "The alias of the destination blog for importing your Hugo content.",
				Required:    true,
				Destination: &dstBlog,
			},

			&cli.StringFlag{
				Name:        "content-dir",
				Usage:       "The name of the path to your source Hugo content (e.g., to import /content/news, use --content-dir news)",
				Required:    true,
				Destination: &srcPath,
			},

			&cli.StringFlag{
				Name:        "instance",
				Aliases:     []string{"i"},
				Usage:       "Provide the URL of your WriteFreely instance (e.g., '--instance https://pencil.writefree.ly')",
				Destination: &instanceUrl,
			},

			&cli.BoolFlag{
				Name:        "images",
				Usage:       "Use this flag to import local images to Snap.as (only for Write.as accounts with Snap.as add-on).",
				Destination: &uploadImages,
			},
		},

		Action: func(c *cli.Context) error {
			if uploadImages && len(instanceUrl) > 0 {
				fmt.Println("Uploading images to Snap.as is only available on Write.as!")
				return nil
			}

			fmt.Println("Hello", username)

			fmt.Println("Please enter your password:")

			var enteredPassword string
			for {
				password, err := term.ReadPassword(0)
				if err != nil {
					panic(err)
				}
				if len(password) != 0 {
					fmt.Println("Press Return to log in and start the migration.")
					enteredPassword = string(password)
				} else {
					break
				}
			}
			err := SignIn(username, enteredPassword, instanceUrl)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Importing content from content ->", srcPath)
			fmt.Println("Importing content into blog alias ->", dstBlog)
			if uploadImages {
				fmt.Println("Uploading local images to Snap.as")
			}
			posts, err := ParseContentDirectory(srcPath, uploadImages)
			if err != nil {
				SignOut()
				log.Fatal(err)
			}
			for _, post := range posts {
				err := PublishPost(post, dstBlog)
				if err != nil {
					SignOut()
					log.Fatal(err)
				}
			}
			err = WriteResponsesToDisk()
			if err != nil {
				SignOut()
				log.Fatal(err)
			}

			SignOut()

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
