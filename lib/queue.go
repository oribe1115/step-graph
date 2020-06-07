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

func (q *Queue) Sprint() string {
	result := ""
	for _, n := range q.Nodes {
		result += fmt.Sprintf("{%d: %s} ", n.ID, n.Name)
	}
	return result
}

func (q *Queue) Enqueue(node *Node) {
	q.Nodes = append(q.Nodes, node)
}

func (q *Queue) Dequeue() (*Node, error) {
	if q.Len() == 0 {
		return nil, fmt.Errorf("faild to dequeue")
	}

	node := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return node, nil
}
