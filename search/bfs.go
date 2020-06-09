package search

import (
	"fmt"

	"github.com/oribe1115/step-graph/lib"
)

func BreadthFirstSearch(graph *lib.Graph, startName string, targetName string) (target *lib.Node, depth int, route []*lib.Node, err error) {
	searchRecord := lib.InitSearchRecord()
	queue := lib.InitQueue()

	startNode := graph.FindNodeByName(startName)
	if startNode == nil {
		return nil, 0, nil, fmt.Errorf("faild to found startNode. startName=%s", startName)
	}

	targetNode := graph.FindNodeByName(targetName)
	if targetNode == nil {
		return nil, 0, nil, fmt.Errorf("faild to found targetNode. targetName=%s", targetName)
	}

	searchRecord.AddRecord(startNode, 0, nil)
	queue.Enqueue(startNode)

	now := &lib.Node{}
	nowRecord := &lib.Record{}

	for queue.Len() != 0 {
		now, err = queue.Dequeue()
		if err != nil {
			return nil, 0, nil, err
		}

		nowRecord = searchRecord.GetRecord(now.ID)

		if now.Name == targetName {
			break
		}

		for _, node := range now.Links {
			if searchRecord.IsRecorded(node.ID) {
				continue
			}

			queue.Enqueue(node)
			searchRecord.AddRecord(node, nowRecord.Count+1, now)
		}
	}

	route, err = searchRecord.GetRoute(now)
	if err != nil {
		return now, nowRecord.Count, nil, err
	}

	return now, nowRecord.Count, route, nil
}

func FindFarthermostNode(graph *lib.Graph, startName string) (end *lib.Node, depth int, route []*lib.Node, err error) {
	searchRecord := lib.InitSearchRecord()
	queue := lib.InitQueue()

	startNode := graph.FindNodeByName(startName)
	if startNode == nil {
		return nil, 0, nil, fmt.Errorf("faild to found startNode. startName=%s", startName)
	}

	searchRecord.AddRecord(startNode, 0, nil)
	queue.Enqueue(startNode)

	now := &lib.Node{}
	nowRecord := &lib.Record{}

	for queue.Len() != 0 {
		now, err = queue.Dequeue()
		if err != nil {
			return nil, 0, nil, err
		}

		nowRecord = searchRecord.GetRecord(now.ID)

		for _, node := range now.Links {
			if searchRecord.IsRecorded(node.ID) {
				continue
			}

			queue.Enqueue(node)
			searchRecord.AddRecord(node, nowRecord.Count+1, now)
		}
	}

	route, err = searchRecord.GetRoute(now)
	if err != nil {
		return now, nowRecord.Count, nil, err
	}

	return now, nowRecord.Count, route, nil
}
