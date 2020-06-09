package lib

import "container/heap"

type PriorityQueue struct {
	pq priorityQueue
}

func InitPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(&pq.pq)
	return pq
}

func (pq *PriorityQueue) Push(node *Node, priority int) {
	item := &pqItem{
		value:    node,
		priority: priority,
	}
	heap.Push(&pq.pq, item)
}

func (pq *PriorityQueue) Pop() (node *Node, priority int) {
	if pq.pq.Len() == 0 {
		return nil, 0
	}
	item := heap.Pop(&pq.pq).(*pqItem)
	return item.value, item.priority
}

type pqItem struct {
	value    *Node
	priority int
	index    int
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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
