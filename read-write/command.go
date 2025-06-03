// TODO come back to this when the main.go is ready
package main

import (
	"flag"
	"fmt"
	"os"
)

type CmdFlags struct {
	Add     bool
	Title   string
	Content string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.BoolVar(&cf.Add, "a", false, "Create a new Note")
	flag.StringVar(&cf.Title, "c", "", "Specify Title of the new note ")
	flag.StringVar(&cf.Title, "c", "", "Specify content of the new note")

	flag.Parse()

	if cf.Add {
		if cf.Title == "" || cf.Content == "" {
			fmt.Println("Title and Content required for creation")
			os.Exit(1)
		}
	}
	return &cf
}
