package lib

import "fmt"

// Queue Nodeのキュー
type Queue struct {
	Nodes []*Node
}

// InitQueue Queueを初期化して返す
func InitQueue() *Queue {
	queue := &Queue{}
	queue.Nodes = make([]*Node, 0)
	return queue
}

// Len Queueの長さを返す
func (q *Queue) Len() int {
	return len(q.Nodes)
}

// Sprint Queueの中身を文字列にして返す
func (q *Queue) Sprint() string {
	result := ""
	for _, n := range q.Nodes {
		result += fmt.Sprintf("%s ", n.Sprint())
	}
	return result
}

// Enqueue enqueueを行う
func (q *Queue) Enqueue(node *Node) {
	q.Nodes = append(q.Nodes, node)
}

// Dequeue dequeueを行う
// Queueが空のときはエラーを返す
func (q *Queue) Dequeue() (*Node, error) {
	if q.Len() == 0 {
		return nil, fmt.Errorf("faild to dequeue")
	}

	node := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return node, nil
}
