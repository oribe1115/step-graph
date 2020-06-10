package lib

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// NodeData ファイルから読み取ったノードの情報
type NodeData struct {
	ID   int
	Name string
}

// LinkData ファイルから読み取ったリンクの情報
type LinkData struct {
	FirstID  int
	SecondID int
}

// EdgeCostData ファイルから読み取った辺のコスト情報
type EdgeCostData struct {
	FirstID  int
	SecondID int
	Cost     int
}

// ReadNodeData 指定したファイルからノードの情報を読み取って返す
func ReadNodeData(filename string) ([]NodeData, error) {
	data, err := readData(filename)
	if err != nil {
		return nil, err
	}
	nodeDataList := make([]NodeData, 0)
	for _, d := range data {
		if d == "" {
			continue
		}
		tmp := strings.Split(d, "\t")
		id, err := strconv.Atoi(tmp[0])
		if err != nil {
			return nodeDataList, err
		}
		nodeDataList = append(nodeDataList, NodeData{ID: id, Name: tmp[1]})
	}

	return nodeDataList, nil
}

// ReadLinkData 指定したファイルからリンクの情報を読み取って返す
func ReadLinkData(filename string) ([]LinkData, error) {
	data, err := readData(filename)
	if err != nil {
		return nil, err
	}
	linkDataList := make([]LinkData, 0)
	for _, d := range data {
		if d == "" {
			continue
		}
		tmp := strings.Split(d, "\t")
		firstID, err := strconv.Atoi(tmp[0])
		if err != nil {
			return linkDataList, err
		}
		secondID, err := strconv.Atoi(tmp[1])
		if err != nil {
			return linkDataList, err
		}
		linkDataList = append(linkDataList, LinkData{FirstID: firstID, SecondID: secondID})
	}

	return linkDataList, nil
}

// ReadEdgeCostData 指定したファイルから辺のコスト情報を読み取って返す
func ReadEdgeCostData(filename string) ([]EdgeCostData, error) {
	data, err := readData(filename)
	if err != nil {
		return nil, err
	}

	edgeCostDataList := make([]EdgeCostData, 0)

	for _, d := range data {
		if d == "" {
			continue
		}
		tmp := strings.Split(d, "\t")

		firstID, err := strconv.Atoi(tmp[0])
		if err != nil {
			return edgeCostDataList, err
		}

		secondID, err := strconv.Atoi(tmp[1])
		if err != nil {
			return edgeCostDataList, err
		}

		cost, err := strconv.Atoi(tmp[2])
		if err != nil {
			return edgeCostDataList, err
		}

		edgeCostDataList = append(edgeCostDataList, EdgeCostData{firstID, secondID, cost})
	}

	return edgeCostDataList, nil
}

func readData(filename string) ([]string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	data := strings.Split(string(b), "\n")
	return data, nil
}
