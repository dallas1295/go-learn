package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{Filename: fileName}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error saving to storage:", err)
		return err
	}
	// 0644, Owner to read and write to the file all others can read the file
	return os.WriteFile(s.Filename, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.Filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	return json.Unmarshal(fileData, data)
}
