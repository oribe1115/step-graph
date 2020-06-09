package topic

import "github.com/oribe1115/step-graph/lib"

type Stations struct {
	Graph    *lib.Graph
	EdgeCost *lib.EdgeCost
}

func InitStations() (*Stations, error) {
	stations := &Stations{}
	stations.Graph = lib.InitGraph()
	stations.EdgeCost = lib.InitEdgeCost(false)

	nodeData, err := lib.ReadNodeData("./data/stations/stations.txt")
	if err != nil {
		return nil, err
	}

	for _, nd := range nodeData {
		node, err := lib.CreateNode(nd.ID, nd.Name)
		if err != nil {
			return nil, err
		}

		err = stations.Graph.SetNode(node)
		if err != nil {
			return nil, err
		}
	}

	edgeCostData, err := lib.ReadEdgeCostData("./data/stations/stations.txt")
	if err != nil {
		return nil, err
	}

	for _, ed := range edgeCostData {
		err := stations.Graph.LinkNodes(ed.FirstID, ed.SecondID)
		if err != nil {
			return nil, err
		}

		err = stations.Graph.LinkNodes(ed.SecondID, ed.FirstID)
		if err != nil {
			return nil, err
		}

		stations.EdgeCost.Set(ed.FirstID, ed.SecondID, ed.Cost)
	}

	return stations, nil
}
