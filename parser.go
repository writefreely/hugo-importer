package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

var cmdParseSource = cli.Command{
	Name: "parse",
	Usage: "Parse the content source directory",
	Action: parseContentDirectory,
}

func parseContentDirectory(c *cli.Context) error {
	var numberOfFiles int = 0

	wd, err := os.Getwd()
	fmt.Println("Working Directory ->", wd)

	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(wd, func(path string, i os.FileInfo, err error) error {
		if !i.IsDir() && !strings.HasPrefix(i.Name(), ".") && strings.HasSuffix(i.Name(), ".go") {
			fileName := i.Name()
			fmt.Println("> Parsing", fileName)
			parsePost(fileName)
			fmt.Println("> Finished parsing", fileName)
			numberOfFiles += 1
		}
		return nil
	})

	fmt.Printf("Parsed %d files", numberOfFiles)

	return nil
}

func parsePost(f string) error {
	file, err := os.Open(f)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// For now, just print the output to the console.
		fmt.Println(scanner.Text())
	}

	return nil
}