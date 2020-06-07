package lib

import "fmt"

type Graph struct {
	Nodes map[int]*Node
}

func InitGraph() *Graph {
	return &Graph{}
}

func (g *Graph) FindNode(id int) *Node {
	node, ok := g.Nodes[id]
	if !ok {
		return nil
	}

	return node
}

func (g *Graph) SetNode(node *Node) error {
	if g.FindNode(node.ID) != nil {
		return fmt.Errorf("deplicate id. id=%d", node.ID)
	}
	g.Nodes[node.ID] = node
	return nil
}

func (g *Graph) LinkNodes(fromID int, toID int) error {
	from := g.FindNode(fromID)
	if from == nil {
		return fmt.Errorf("not found fromNode. fromID=%d", fromID)
	}
	to := g.FindNode(toID)
	if to == nil {
		return fmt.Errorf("not found toNode. toID=%d", toID)
	}

	from.LinkTo(to)
	return nil
}
