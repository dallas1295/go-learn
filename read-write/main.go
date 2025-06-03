package main

import (
	"bufio"
	"fmt"
	"os"
	"read-write/note"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter note name:")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid string input:", err)
		os.Exit(1)
	}
	fmt.Print("Enter note content:")
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid string input:", err)
		os.Exit(1)
	}

	note.Create(title, content)
	if err != nil {
		fmt.Println("Error creating note")
	}

	fmt.Println("Note has been created")

	fmt.Print("Enter note name for search:")
	title, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Note does not exist")
	}

	note.Read(title)

	fmt.Println("Enter file to be deleted:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Note does not exist")
	}
	fmt.Println("are you sure? type y to proceed")
	confirm, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response")
	}

	confirm = strings.TrimSpace(confirm)

	if confirm != "y" {
		fmt.Println("Invalid input exiting application")
	} else if confirm == "y" {
		note.Delete(name, true)
	} else {
		note.Delete(name, false)
	}

	fmt.Println("Select note to change name:")
	oldName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Note does not exist")
	}
	fmt.Println("What would you like to rename it to?")
	newName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid naming structure")
	}

	note.EditTitle(oldName, newName)
}
