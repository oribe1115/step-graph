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

	targetNode := graph.FindNodeByName(targetName)
	if targetNode == nil {
		return nil, 0, nil, fmt.Errorf("faild to found targetNode. targetName=%s", targetName)
	}

	priorityQueue.Push(startNode, nil, 0)

	now := &lib.Node{}
	from := &lib.Node{}
	tmpCost := 0

	for priorityQueue.Len() != 0 {
		now, from, tmpCost = priorityQueue.Pop()
		if now.Name == targetName {
			break
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

	searchRecord.AddRecord(now, tmpCost, from)
	route, err = searchRecord.GetRoute(now)
	if err != nil {
		return now, tmpCost, nil, err
	}

	return now, tmpCost, route, nil
}

func JustCostRouteWithDijkstra(graph *lib.Graph, startName string, targetCost int, getCost func(from *lib.Node, to *lib.Node) (int, error)) (routes [][]*lib.Node, err error) {
	searchRecord := lib.InitSearchRecord()
	priorityQueue := lib.InitPriorityQueue()

	startNode := graph.FindNodeByName(startName)
	if startNode == nil {
		return nil, fmt.Errorf("faild to found startNode. startName=%s", startName)
	}

	priorityQueue.Push(startNode, nil, 0)

	now := &lib.Node{}
	from := &lib.Node{}
	tmpCost := 0
	routes = make([][]*lib.Node, 0)

	for priorityQueue.Len() != 0 {
		now, from, tmpCost = priorityQueue.Pop()
		if tmpCost > targetCost {
			break
		}

		if tmpCost == targetCost {
			searchRecord.AddRecord(now, tmpCost, from)
			route, _ := searchRecord.GetRoute(now)
			routes = append(routes, route)
			continue
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
				return nil, err
			}
			priorityQueue.Push(node, now, tmpCost+newCost)
		}
	}

	return routes, nil
}
