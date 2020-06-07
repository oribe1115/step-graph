package search

import (
	"fmt"

	"github.com/oribe1115/step-graph/lib"
)

func BreadthFirstSearch(graph *lib.Graph, startName string, targetName string) (target *lib.Node, depth int, err error) {
	searchRecord := lib.InitSearchRecord()
	queue := lib.InitQueue()

	startNode := graph.FindNodeByName(startName)
	if startNode == nil {
		return nil, 0, fmt.Errorf("faild to found startNode. startName=%s", startName)
	}

	searchRecord.AddRecord(startNode.ID, 0)
	queue.Enqueue(startNode)

	for queue.Len() != 0 {
		now, err := queue.Dequeue()
		if err != nil {
			return nil, 0, err
		}

		nowRecord := searchRecord.GetRecord(now.ID)

		if now.Name == targetName {
			return now, nowRecord.Count, nil
		}

		for _, node := range now.Links {
			if searchRecord.IsRecorded(node.ID) {
				continue
			}

			queue.Enqueue(node)
			searchRecord.AddRecord(node.ID, nowRecord.Count+1)
		}
	}

	return nil, 0, nil
}
