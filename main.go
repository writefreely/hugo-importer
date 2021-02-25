package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var username string
	var dstBlog string
	var srcPath string
	var uploadImages bool

	app := &cli.App{
		Name: "Write.as Hugo Importer",
		Usage: "Import a Hugo source directory into Write.as/WriteFreely by running this importer from the /content directory of your Hugo site.",

		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "user",
				Aliases: []string{"u"},
				Usage: "The username for the Write.as/WriteFreely account",
				Required: true,
				Destination: &username,
			},

			&cli.StringFlag{
				Name: "blog",
				Aliases: []string{"b"},
				Usage: "The alias of the destination blog for importing your Hugo content.",
				Required: true,
				Destination: &dstBlog,
			},

			&cli.StringFlag{
				Name: "content-dir",
				Usage: "The name of the path to your source Hugo content (e.g., to import /content/news, use --content-dir news)",
				Required: true,
				Destination: &srcPath,
			},

			&cli.BoolFlag{
				Name: "images",
				Usage: "Use this flag to import local images to Snap.as (only for Write.as accounts with Snap.as add-on).",
				Destination: &uploadImages,
			},
		},

		Action: func(c *cli.Context) error {
			fmt.Println("Hello", username)
			fmt.Println("Importing content from content ->", srcPath)
			fmt.Println("Importing content into blog alias ->", dstBlog)
			if uploadImages {
				fmt.Println("> Uploading local images to Snap.as")
			}
			ParseContentDirectory(srcPath)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}