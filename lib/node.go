package lib

import "fmt"

type Node struct {
	ID    int
	Name  string
	Links []*Node
}

func CreateNode(id int, name string) (*Node, error) {
	if name == "" {
		return nil, fmt.Errorf("name is blank")
	}
	return &Node{
		ID:   id,
		Name: name,
	}, nil
}

func (n *Node) LinkTo(to *Node) {
	n.Links = append(n.Links, to)
}

func (n *Node) Sprint() string {
	return fmt.Sprintf("{%d: %s}", n.ID, n.Name)
}

// SprintNodeListAsRoute "{ID: Name} -> {ID: Name}"の形の文字列にする
func SprintNodeListAsRoute(nodeList []*Node) string {
	result := ""
	for _, node := range nodeList {
		result += fmt.Sprintf("{%d: %s} -> ", node.ID, node.Name)
	}
	result = result[:len(result)-4]
	return result
}
