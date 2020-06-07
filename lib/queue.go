package lib

import "fmt"

type Queue struct {
	Nodes []*Node
}

func InitQueue() *Queue {
	queue := &Queue{}
	queue.Nodes = make([]*Node, 0)
	return queue
}

func (q *Queue) Len() int {
	return len(q.Nodes)
}

func (q *Queue) Print() {
	for _, n := range q.Nodes {
		fmt.Printf("%v", n)
	}
	fmt.Printf("\n")
}

func (q *Queue) Enqueue(node *Node) {
	q.Nodes = append(q.Nodes, node)
}

func (q *Queue) Dequeue() *Node {
	if q.Len() == 0 {
		return nil
	}

	node := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return node
}
