package lib

import "fmt"

// Graph Nodeの情報を連想配列で保持したグラフ
type Graph struct {
	Nodes map[int]*Node
}

// InitGraph Graphを初期化して返す
func InitGraph() *Graph {
	return &Graph{
		Nodes: map[int]*Node{},
	}
}

// FindNodeByID グラフから該当するidのNodeを返す
func (g *Graph) FindNodeByID(id int) *Node {
	node, ok := g.Nodes[id]
	if !ok {
		return nil
	}

	return node
}

// FindNodeByName GraphからNameが一致しているNodeを返す
// グラフ内でNodeの重複が存在していない前提
func (g *Graph) FindNodeByName(name string) *Node {
	for _, node := range g.Nodes {
		if node.Name == name {
			return node
		}
	}

	return nil
}

// SetNode GraphにNodeを追加する
func (g *Graph) SetNode(node *Node) error {
	if g.FindNodeByID(node.ID) != nil {
		return fmt.Errorf("deplicate id. id=%d", node.ID)
	}
	g.Nodes[node.ID] = node
	return nil
}

// LinkNodes Graph内のNodeにリンク情報を追加する
func (g *Graph) LinkNodes(fromID int, toID int) error {
	from := g.FindNodeByID(fromID)
	if from == nil {
		return fmt.Errorf("not found fromNode. fromID=%d", fromID)
	}
	to := g.FindNodeByID(toID)
	if to == nil {
		return fmt.Errorf("not found toNode. toID=%d", toID)
	}

	from.LinkTo(to)
	return nil
}
