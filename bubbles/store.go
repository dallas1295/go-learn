package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

type Note struct {
	ID    int64
	Title string
	Body  string
}

type Store struct {
	notesDir string
}

func (s *Store) Init() error {
	// Get the user's home directory
	usr, err := user.Current()
	if err != nil {
		return err
	}

	// Set the notes directory path
	s.notesDir = filepath.Join(usr.HomeDir, "Documents", "notes")

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(s.notesDir, 0755); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetNotes() ([]Note, error) {
	entries, err := os.ReadDir(s.notesDir)
	if err != nil {
		return nil, err
	}

	var notes []Note

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(s.notesDir, entry.Name())
		file, err := os.Open(path)
		if err != nil {
			continue // skip files that can't be opened
		}

		scanner := bufio.NewScanner(file)
		var note Note
		var bodyLines []string

		// Parse ID
		if scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "ID: ") {
				note.ID, _ = strconv.ParseInt(strings.TrimPrefix(line, "ID: "), 10, 64)
			}
		}
		// Parse Title
		if scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "Title: ") {
				note.Title = strings.TrimPrefix(line, "Title: ")
			}
		}
		// Skip blank line
		scanner.Scan()

		// The rest is the body
		for scanner.Scan() {
			bodyLines = append(bodyLines, scanner.Text())
		}
		note.Body = strings.Join(bodyLines, "\n")
		notes = append(notes, note)

		file.Close()
	}

	return notes, nil
}

func (s *Store) nextID() (int64, error) {
	notes, err := s.GetNotes()
	if err != nil {
		return 1, nil // If error, start at 1
	}
	var maxID int64
	for _, n := range notes {
		if n.ID > maxID {
			maxID = n.ID
		}
	}
	return maxID + 1, nil
}

func (s *Store) SaveNote(note Note) error {
	if note.ID == 0 {
		id, err := s.nextID()
		if err != nil {
			return err
		}
		note.ID = id
	}

	filename := fmt.Sprintf("note-%d.txt", note.ID)
	path := filepath.Join(s.notesDir, filename)
	content := fmt.Sprintf("ID: %d\nTitle: %s\n\n%s", note.ID, note.Title, note.Body)
	return os.WriteFile(path, []byte(content), 0644)
}
