package main

import "fmt"

func main() {
	s := 10

	for i := 1; i <= s; i++ {
		if i >= 4 {
			fmt.Println("Blub Blub")
		}
		fmt.Println("have an,", i)
		for j := i; j >= 2; j-- {
			fmt.Println("it's workings!!!")
		}
	}
}
