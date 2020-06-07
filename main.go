package main

import (
	"fmt"

	"github.com/oribe1115/step-graph/lib"
)

func main() {
	queue := lib.InitQueue()
	queue.Print()

	queue.Enqueue(lib.CreateNode(0, "a"))
	fmt.Println(queue.Len())
	queue.Print()
	queue.Enqueue(lib.CreateNode(1, "b"))
	fmt.Println(queue.Len())
	queue.Print()

	n := queue.Dequeue()
	fmt.Println(n)
	queue.Print()
	n = queue.Dequeue()
	fmt.Println(n)
	queue.Print()
	n = queue.Dequeue()
	fmt.Println(n)
	queue.Print()
}
