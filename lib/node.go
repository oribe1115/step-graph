package lib

import "fmt"

// Node グラフの頂点情報とそこからのびる辺の情報を格納したもの
type Node struct {
	ID    int
	Name  string
	Links []*Node
}

// CreateNode Nodeを作成して返す
func CreateNode(id int, name string) (*Node, error) {
	if name == "" {
		return nil, fmt.Errorf("name is blank")
	}
	return &Node{
		ID:   id,
		Name: name,
	}, nil
}

// LinkTo Nodeに辺の情報を追加
func (n *Node) LinkTo(to *Node) {
	n.Links = append(n.Links, to)
}

// Sprint Nodeの情報を文字列にして返す
func (n *Node) Sprint() string {
	return fmt.Sprintf("{%d: %s}", n.ID, n.Name)
}

// SprintNodeListAsRoute "{ID: Name} -> {ID: Name}"の形の文字列にする
func SprintNodeListAsRoute(nodeList []*Node) string {
	result := ""
	for _, node := range nodeList {
		result += fmt.Sprintf("%s -> ", node.Sprint())
	}
	result = result[:len(result)-4]
	return result
}
