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

	searchRecord.AddRecord(startNode, 0, nil)
	queue.Enqueue(startNode)

	for queue.Len() != 0 {
		now, err := queue.Dequeue()
		if err != nil {
			return nil, 0, nil, err
		}

		nowRecord := searchRecord.GetRecord(now.ID)

		if now.Name == targetName {
			route, err := searchRecord.GetRoute(now)
			if err != nil {
				return now, nowRecord.Count, nil, err
			}
			return now, nowRecord.Count, route, nil
		}

		for _, node := range now.Links {
			if searchRecord.IsRecorded(node.ID) {
				continue
			}

			queue.Enqueue(node)
			searchRecord.AddRecord(node, nowRecord.Count+1, now)
		}
	}

	return nil, 0, nil, nil
}
