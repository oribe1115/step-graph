package lib

import "fmt"

type Record struct {
	Node  *Node
	Count int   // 探索の深さなどの記録に使用
	From  *Node // どのノードからこのIDのノードに来たのか
}

type SearchRecord struct {
	Records map[int]*Record
}

func InitSearchRecord() *SearchRecord {
	return &SearchRecord{
		Records: map[int]*Record{},
	}
}

func (s *SearchRecord) IsRecorded(id int) bool {
	_, ok := s.Records[id]
	return ok
}

func (s *SearchRecord) AddRecord(node *Node, count int, from *Node) {
	record := &Record{
		Node:  node,
		Count: count,
		From:  from,
	}
	s.Records[node.ID] = record
}

func (s *SearchRecord) GetRecord(id int) *Record {
	record, ok := s.Records[id]
	if !ok {
		return nil
	}
	return record
}

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
