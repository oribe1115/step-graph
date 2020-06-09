package lib

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type NodeData struct {
	ID   int
	Name string
}

type LinkData struct {
	FirstID  int
	SecondID int
}

type EdgeCostData struct {
	FirstID  int
	SecondID int
	Cost     int
}

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
