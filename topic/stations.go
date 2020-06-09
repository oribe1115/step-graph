package topic

import (
	"fmt"

	"github.com/oribe1115/step-graph/lib"
	"github.com/oribe1115/step-graph/search"
)

type Stations struct {
	Graph    *lib.Graph
	EdgeCost *lib.EdgeCost
}

func CmdStaions() {
	stations, err := InitStations()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Select mode with Staions")
	fmt.Println("1. Breadth First Search")
	fmt.Printf("> ")
	input := lib.ReadLine()

	switch input {
	case "1":
		fmt.Println("Input start name")
		fmt.Printf("> ")
		startName := lib.ReadLine()

		fmt.Println("Input target name")
		fmt.Printf("> ")
		targetName := lib.ReadLine()

		target, depth, route, err := stations.BreadthFirstSearch(startName, targetName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("target: %s, depth: %d\n", target.Sprint(), depth)
		fmt.Printf("route: %s\n", lib.SprintNodeListAsRoute(route))
	default:
		fmt.Println("Invalid input")
	}
}

func InitStations() (*Stations, error) {
	stations := &Stations{}
	stations.Graph = lib.InitGraph()
	stations.EdgeCost = lib.InitEdgeCost(false)

	nodeData, err := lib.ReadNodeData("./data/stations/stations.txt")
	if err != nil {
		return nil, err
	}

	for _, nd := range nodeData {
		node, err := lib.CreateNode(nd.ID, nd.Name)
		if err != nil {
			return nil, err
		}

		err = stations.Graph.SetNode(node)
		if err != nil {
			return nil, err
		}
	}

	edgeCostData, err := lib.ReadEdgeCostData("./data/stations/edges.txt")
	if err != nil {
		return nil, err
	}

	for _, ed := range edgeCostData {
		err := stations.Graph.LinkNodes(ed.FirstID, ed.SecondID)
		if err != nil {
			return nil, err
		}

		err = stations.Graph.LinkNodes(ed.SecondID, ed.FirstID)
		if err != nil {
			return nil, err
		}

		stations.EdgeCost.Set(ed.FirstID, ed.SecondID, ed.Cost)
	}

	return stations, nil
}

func (s *Stations) BreadthFirstSearch(startName string, targetName string) (target *lib.Node, depth int, route []*lib.Node, err error) {
	return search.BreadthFirstSearch(s.Graph, startName, targetName)
}
