package main

import (
	"fmt"
	"log/slog"
	"os"
)

type person struct {
	name string
	age  int
}

const (
	// \033 is an octal (stands for escape)
	// [ control sequence introducer for arguments
	// 0m is select Graphic Rendition, 0 being a reset all attribs, m denoting the end of the SGR
	// next comes color

	/*
		Text attributes (bold, faint, italic, underline, blink, inverse, hidden, strikethrough) 1-9
		Foreground colors (30-37, 90-97 for bright)
		Background colors (40-47, 100-107 for bright)
	*/

	Reset   = "\033[0m" // reset attributes
	Red     = "\033[31m"
	Bold    = "\033[1m"
	Italics = "\033[2m"
)

var ErrUnder18 = fmt.Errorf("under 18 err")

func validateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age: %d (%w)", age, ErrUnder18)
	}
	return nil
}

func main() {
	name := "dallas"
	age := 29 // if age below 18 the validateAge function with throw error
	information := fmt.Sprintf("Name: %s, Age: %d", name, age)

	// standard print line (adds space after params)
	fmt.Println("Name:", name, "Age:", age)

	// printf asserting variables in the quotes, does NOT put a new line
	// used for custom formatting
	/*
	 %s -> string
	 %d -> integer
	 %f -> float (can truncate like %.2f)
	 %t -> boolean
	 %v -> default formatting (like Println w/o new line)
	 %+v -> struct
	 %#v -> go syntax representation
	*/

	fmt.Printf("Name: %s Age: %d\n", name, age)
	s := []int{1, 2, 3}
	p := person{"dallas", 29}
	fmt.Printf("struct: %+v\n", p)
	fmt.Printf("slice: %#v\n", s)

	// time for colors
	fmt.Println(Red + "This is bold red text" + Reset)
	// can also put it inside the quotes WITHOUT const object
	fmt.Println("\033[33;1;3;4mThis is a yellow line\033[0m")

	fmt.Println(information)

	if err := validateAge(age); err != nil {
		fmt.Println(err)
	}

	// Logging Stuff
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Debug Message")
	logger.Info("Info Message")
	logger.Warn("Warn Message")
	logger.Error("Error Message")
}
