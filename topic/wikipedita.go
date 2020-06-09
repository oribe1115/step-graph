package topic

import (
	"fmt"

	"github.com/oribe1115/step-graph/lib"
	"github.com/oribe1115/step-graph/search"
)

type Wikipedia struct {
	Graph *lib.Graph
}

func CmdWikipedia() {
	fmt.Println("Initializing Wikipedia data...")
	wikipedia, err := InitWikipedia()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Select mode with Wikipedia")
	fmt.Println("1. Search Route by Breadth First Search")
	fmt.Println("2. Find Farthermost Node")
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

		target, depth, route, err := wikipedia.BreadthFirstSearch(startName, targetName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("target: %s, depth: %d\n", target.Sprint(), depth)
		fmt.Printf("route: %s\n", lib.SprintNodeListAsRoute(route))
		return
	case "2":
		fmt.Println("Input start name")
		fmt.Printf("> ")
		startName := lib.ReadLine()

		end, depth, route, err := wikipedia.FindFathermostNode(startName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("end: %s, depth: %d\n", end.Sprint(), depth)
		fmt.Printf("route: %s\n", lib.SprintNodeListAsRoute(route))
		return
	default:
		fmt.Println("Invalid input")
	}
}

func InitWikipedia() (*Wikipedia, error) {
	wikipedia := &Wikipedia{}
	wikipedia.Graph = lib.InitGraph()

	nodeData, err := lib.ReadNodeData("./data/wikipedia/pages.txt")
	if err != nil {
		return nil, err
	}

	for _, nd := range nodeData {
		node, err := lib.CreateNode(nd.ID, nd.Name)
		if err != nil {
			return nil, err
		}

		err = wikipedia.Graph.SetNode(node)
		if err != nil {
			return nil, err
		}
	}

	linkData, err := lib.ReadLinkData("./data/wikipedia/links.txt")
	if err != nil {
		return nil, err
	}

	for _, ld := range linkData {
		err = wikipedia.Graph.LinkNodes(ld.FirstID, ld.SecondID)
		if err != nil {
			return nil, err
		}
	}

	return wikipedia, nil
}

func (w *Wikipedia) BreadthFirstSearch(startName string, targetName string) (target *lib.Node, depth int, route []*lib.Node, err error) {
	return search.BreadthFirstSearch(w.Graph, startName, targetName)
}

func (w *Wikipedia) FindFathermostNode(startName string) (end *lib.Node, depth int, route []*lib.Node, err error) {
	startNode := w.Graph.FindNodeByName(startName)
	if startNode == nil {
		return nil, 0, nil, fmt.Errorf("faild to found startNode. startName=%s", startName)
	}

	return search.FindFarthermostNode(w.Graph, startName)
}
