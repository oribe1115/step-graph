package main

import (
	"fmt"

	"github.com/oribe1115/step-graph/lib"
	"github.com/oribe1115/step-graph/topic"
)

func main() {
	lib.InitStdin()

	fmt.Println("Select topic")
	fmt.Println("1. SNS")
	fmt.Println("2. Stations")
	fmt.Printf("> ")
	input := lib.ReadLine()

	switch input {
	case "1":
		topic.CmdSns()
		return
	case "2":
		topic.CmdStaions()
		return
	default:
		fmt.Println("Invalid input")
	}
}
