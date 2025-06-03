package note

import (
	"fmt"
	"os"
	"strings"
)

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Color = "\033[34m"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func titleToFilename(title string) string {
	trimmed := strings.TrimSpace(title)
	filename := trimmed + ".txt"

	return filename
}

func Create(title string, content string) {
	filename := titleToFilename(title)
	data := []byte(content)

	if fileExists(filename) {
		fmt.Println("That file already exists!")
		return
	}

	os.WriteFile(filename, data, 0666)
}

func Read(title string) {
	filename := titleToFilename(title)
	if !fileExists(filename) {
		fmt.Println("This file does not exist")
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Failed to read note contents")
	}

	content := string(data)

	fmt.Println("Note:", Color+Bold+title+Reset)
	fmt.Printf("%s", content)
}

func Delete(title string, confirm bool) {
	if !confirm {
		fmt.Println("Delete cancelled")
		return
	}

	filename := titleToFilename(title)

	if !fileExists(filename) {
		fmt.Println("File does not exist")
		return
	}

	err := os.Remove(filename)
	if err != nil {
		fmt.Println("File could not be deleted")
	}
	fmt.Println("File Deleted")
}

func EditTitle(oldTitle string, newTitle string) {
	oldFilename := titleToFilename(oldTitle)
	if !fileExists(oldFilename) {
		fmt.Println("File does not exist")
		return
	}

	newTrimmed := strings.TrimSpace(newTitle)
	newFilename := titleToFilename(newTitle)

	if fileExists(newFilename) {
		fmt.Println("the file", Bold+Color+newTrimmed, "already exists")
		return
	}

	err := os.Rename(oldFilename, newFilename)
	if err != nil {
		fmt.Println("could not rename file")
		return
	}
	fmt.Println("File name changed to:", newTrimmed)
}
