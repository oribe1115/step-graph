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

	fmt.Println("Input start name")
	fmt.Printf("> ")
	startName := lib.ReadLine()

	fmt.Println("Input target name")
	fmt.Printf("> ")
	targetName := lib.ReadLine()

	target, depth, err := sns.BreadthFirstSearch(startName, targetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("target: {%d: %s}, depth: %d\n", target.ID, target.Name, depth)
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

func (s *Sns) BreadthFirstSearch(startName string, targetName string) (target *lib.Node, depth int, err error) {
	return search.BreadthFirstSearch(s.Graph, startName, targetName)
}
