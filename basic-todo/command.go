// Package main is to run the app
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "a", "", "Add new todo specify title")
	flag.StringVar(&cf.Add, "e", "", "edit todo by index and specify new title, id:new_title")
	flag.IntVar(&cf.Delete, "d", -1, "Specify todo by index to delete")
	flag.IntVar(&cf.Toggle, "t", -1, "Specify todo by index to toggle")
	flag.BoolVar(&cf.List, "l", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit; id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
		}
		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Delete != -1:
		todos.delete(cf.Delete)
	default:
		todos.print()
	}
}
