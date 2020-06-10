package lib

import "fmt"

// Record 通過したノードの情報
type Record struct {
	Node  *Node
	Count int   // 探索の深さなどの記録に使用
	From  *Node // どのノードからこのIDのノードに来たのか
}

// SearchRecord 探索で通過したノードの情報を記録する
type SearchRecord struct {
	Records map[int]*Record
}

// InitSearchRecord 初期化したSearchRecordを返す
func InitSearchRecord() *SearchRecord {
	return &SearchRecord{
		Records: map[int]*Record{},
	}
}

// IsRecorded すでにSearchRecordに記録されているidかどうかを調べる
func (s *SearchRecord) IsRecorded(id int) bool {
	_, ok := s.Records[id]
	return ok
}

// AddRecord SearchRecordに記録を追加する
func (s *SearchRecord) AddRecord(node *Node, count int, from *Node) {
	record := &Record{
		Node:  node,
		Count: count,
		From:  from,
	}
	s.Records[node.ID] = record
}

// GetRecord SearchRecordから指定したidのRecordを返す
func (s *SearchRecord) GetRecord(id int) *Record {
	record, ok := s.Records[id]
	if !ok {
		return nil
	}
	return record
}

// GetRoute 引数が探索の終端のとき、探索の始点から終点までの辿ってきたNodeのリストを返す
func (s *SearchRecord) GetRoute(endNode *Node) ([]*Node, error) {
	route := make([]*Node, 0)

	tmpNode := endNode

	// 無限ループ対策
	limit := len(s.Records)
	count := 0

	for tmpNode != nil {
		if count > limit {
			return route, fmt.Errorf("faild to get correct route")
		}

		route = append([]*Node{tmpNode}, route...)

		tmpRecord := s.GetRecord(tmpNode.ID)
		if tmpRecord == nil {
			return route, fmt.Errorf("faild to find record. id=%d", tmpNode.ID)
		}

		tmpNode = tmpRecord.From
		count++
	}

	return route, nil
}
