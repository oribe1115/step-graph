package lib

import "container/heap"

// PriorityQueue 優先度付きキュー
// heapパッケージを使って実装したpriorityQueueをラップしたもの
type PriorityQueue struct {
	pq priorityQueue
}

// InitPriorityQueue 初期化したPriorityQueueを返す
func InitPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(&pq.pq)
	return pq
}

// Push PriorityQueueにpushする
func (pq *PriorityQueue) Push(node *Node, from *Node, priority int) {
	item := &pqItem{
		node:     node,
		from:     from,
		priority: priority,
	}
	heap.Push(&pq.pq, item)
}

// Pop PriorityQueueから最もpriorityが低いものの情報を返す
// 本来の優先度付きキューの仕様とは逆
func (pq *PriorityQueue) Pop() (node *Node, from *Node, priority int) {
	if pq.pq.Len() == 0 {
		return nil, nil, 0
	}
	item := heap.Pop(&pq.pq).(*pqItem)
	return item.node, item.from, item.priority
}

// Len PriorityQueueの長さを返す
func (pq *PriorityQueue) Len() int {
	return pq.pq.Len()
}

type pqItem struct {
	node     *Node
	from     *Node
	priority int
	index    int
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
