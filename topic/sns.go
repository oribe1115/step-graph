package topic

import (
	"fmt"

	"github.com/oribe1115/step-graph/lib"
	"github.com/oribe1115/step-graph/search"
)

type Sns struct {
	Graph *lib.Graph
}

func CmdSns() {
	sns, err := InitSns()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Select mode with SNS")
	fmt.Println("1. Breadth First Search")
	fmt.Println("2. Find Farthermost")
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

		target, depth, route, err := sns.BreadthFirstSearch(startName, targetName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("target: %s, depth: %d\n", target.Sprint(), depth)
		fmt.Printf("route: %s\n", lib.SprintNodeListAsRoute(route))
		return
	case "2":
		from, to, depth, route, err := sns.FindFarthermost()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("from: %s, to: %s, depth: %d\n", from.Sprint(), to.Sprint(), depth)
		fmt.Printf("route: %s\n", lib.SprintNodeListAsRoute(route))
		return
	default:
		fmt.Println("Invalid input")
	}

}

func InitSns() (*Sns, error) {
	sns := &Sns{}
	sns.Graph = lib.InitGraph()

	nodeData, err := lib.ReadNodeData("./data/sns/nicknames.txt")
	if err != nil {
		return nil, err
	}

	for _, nd := range nodeData {
		node, err := lib.CreateNode(nd.ID, nd.Name)
		if err != nil {
			return nil, err
		}

		err = sns.Graph.SetNode(node)
		if err != nil {
			return nil, err
		}
	}

	linkData, err := lib.ReadLinkData("./data/sns/links.txt")
	if err != nil {
		return nil, err
	}

	for _, ld := range linkData {
		err = sns.Graph.LinkNodes(ld.FirstID, ld.SecondID)
		if err != nil {
			return nil, err
		}
	}

	return sns, nil
}

func (s *Sns) BreadthFirstSearch(startName string, targetName string) (target *lib.Node, depth int, route []*lib.Node, err error) {
	return search.BreadthFirstSearch(s.Graph, startName, targetName)
}

func (s *Sns) FindFarthermost() (from *lib.Node, to *lib.Node, depth int, route []*lib.Node, err error) {
	maxDepth := 0
	var maxFrom *lib.Node
	var maxTo *lib.Node
	var maxRoute []*lib.Node

	for _, from := range s.Graph.Nodes {
		for _, to := range s.Graph.Nodes {
			if from.ID == to.ID {
				continue
			}

			_, depth, route, err := search.BreadthFirstSearch(s.Graph, from.Name, to.Name)
			if err != nil {
				return nil, nil, 0, nil, err
			}
			if depth > maxDepth {
				maxDepth = depth
				maxFrom = from
				maxTo = to
				maxRoute = route
			}
		}
	}

	return maxFrom, maxTo, maxDepth, maxRoute, nil
}
