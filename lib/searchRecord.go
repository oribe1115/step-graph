package lib

import "fmt"

type Record struct {
	ID    int
	Count int // 探索の深さなどの記録に使用
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

func (s *SearchRecord) AddRecord(id int, count int) error {
	if s.IsRecorded(id) {
		return fmt.Errorf("deplicate id. id=%d", id)
	}

	record := &Record{
		ID:    id,
		Count: count,
	}
	s.Records[id] = record

	return nil
}

func (s *SearchRecord) GetRecord(id int) *Record {
	record, ok := s.Records[id]
	if !ok {
		return nil
	}
	return record
}
