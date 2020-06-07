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
