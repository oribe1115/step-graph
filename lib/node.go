package lib

type Node struct {
	ID    int
	Name  string
	Links []*Node
}

func CreateNode(id int, name string) *Node {
	return &Node{
		ID:   id,
		Name: name,
	}
}
