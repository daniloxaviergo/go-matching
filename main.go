package main

import (
	"fmt"

	"go-matching/test"
)

func main() {
	fmt.Println("Hello World!!")
	c := test.Creature{
		Name: "Sammy the Shark",
	}

	fmt.Println(c.Name)
}
