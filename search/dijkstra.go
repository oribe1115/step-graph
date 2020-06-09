package search

import (
	"fmt"

	"github.com/oribe1115/step-graph/lib"
)

func Dijkstra(graph *lib.Graph, startName string, targetName string, getCost func(from *lib.Node, to *lib.Node) (int, error)) (target *lib.Node, totalCost int, route []*lib.Node, err error) {
	searchRecord := lib.InitSearchRecord()
	priorityQueue := lib.InitPriorityQueue()

	startNode := graph.FindNodeByName(startName)
	if startNode == nil {
		return nil, 0, nil, fmt.Errorf("faild to found startNode. startName=%s", startName)
	}

	priorityQueue.Push(startNode, nil, 0)

	for priorityQueue.Len() != 0 {
		now, from, tmpCost := priorityQueue.Pop()
		if now.Name == targetName {
			searchRecord.AddRecord(now, tmpCost, from)
			route, err := searchRecord.GetRoute(now)
			if err != nil {
				return now, tmpCost, nil, err
			}
			return now, tmpCost, route, nil
		}

		// 最小コストが記録済みかどうか
		if searchRecord.IsRecorded(now.ID) {
			continue
		}

		// 最小コストとして記録
		searchRecord.AddRecord(now, tmpCost, from)

		for _, node := range now.Links {
			if searchRecord.IsRecorded(node.ID) {
				continue
			}
			newCost, err := getCost(now, node)
			if err != nil {
				return nil, 0, nil, err
			}
			priorityQueue.Push(node, now, tmpCost+newCost)
		}
	}

	return nil, 0, nil, nil
}
