package lib

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
