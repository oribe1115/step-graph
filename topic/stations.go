package topic

import (
	"fmt"
	"strconv"

	"github.com/oribe1115/step-graph/lib"
	"github.com/oribe1115/step-graph/search"
)

// Stations .
type Stations struct {
	Graph    *lib.Graph
	EdgeCost *lib.EdgeCost
}

// CmdStaions Staionsの関数をCLIで実行する
func CmdStaions() {
	stations, err := InitStations()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Select mode with Staions")
	fmt.Println("1. Search Route by Breadth First Search")
	fmt.Println("2: Search Shortest Route by Dijkstra")
	fmt.Println("3: Search Routes with Just the Time by Dijkstra")
	fmt.Printf("> ")
	input := lib.ReadLine()

	switch input {
	case "1":
		fmt.Println("Input start name")
		fmt.Printf("> ")
		startName := lib.ReadLine()

		fmt.Println("Input target name")
		fmt.Printf("> ")
		targetName := lib.ReadLine()

		target, depth, route, err := stations.BreadthFirstSearch(startName, targetName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("target: %s, depth: %d\n", target.Sprint(), depth)
		fmt.Printf("route: %s\n", lib.SprintNodeListAsRoute(route))
		return
	case "2":
		fmt.Println("Input start name")
		fmt.Printf("> ")
		startName := lib.ReadLine()

		fmt.Println("Input target name")
		fmt.Printf("> ")
		targetName := lib.ReadLine()

		target, requiredTime, route, err := stations.DijkstraWithRequiredTime(startName, targetName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("traget: %s, requiredTime: %d min\n", target.Sprint(), requiredTime)
		fmt.Printf("route: %s\n", lib.SprintNodeListAsRoute(route))
		return
	case "3":
		fmt.Println("Input start name")
		fmt.Printf("> ")
		startName := lib.ReadLine()

		fmt.Println("Input required time")
		fmt.Printf("> ")
		requiredTimeStr := lib.ReadLine()
		requiredTime, err := strconv.Atoi(requiredTimeStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		routes, err := stations.JustTimeRoutesByDijkstra(startName, requiredTime)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(routes) == 0 {
			fmt.Println("Not found such route")
			return
		}

		for _, route := range routes {
			fmt.Println(lib.SprintNodeListAsRoute(route))
		}
		return
	default:
		fmt.Println("Invalid input")
	}
}

// InitStations Stationsを初期化して返す
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

	edgeCostData, err := lib.ReadEdgeCostData("./data/stations/edges.txt")
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

// BreadthFirstSearch startNameからtargetNameまでの最短経路を探索
func (s *Stations) BreadthFirstSearch(startName string, targetName string) (target *lib.Node, depth int, route []*lib.Node, err error) {
	return search.BreadthFirstSearch(s.Graph, startName, targetName)
}

// DijkstraWithRequiredTime startNameからtargetNameまで最短の時間でで行くルートを探索
func (s *Stations) DijkstraWithRequiredTime(startName string, targetName string) (target *lib.Node, totalCost int, route []*lib.Node, err error) {
	getCost := func(from *lib.Node, to *lib.Node) (int, error) {
		cost, ok := s.EdgeCost.Get(from.ID, to.ID)
		if !ok {
			return 0, fmt.Errorf("faild to get cost. fromID=%d toID=%d", from.ID, to.ID)
		}
		return cost, nil
	}
	return search.Dijkstra(s.Graph, startName, targetName, getCost)
}

// JustTimeRoutesByDijkstra startNameからrequiredTimeちょうどでたどり着ける駅を探索
func (s *Stations) JustTimeRoutesByDijkstra(startName string, requiredTime int) ([][]*lib.Node, error) {
	getCost := func(from *lib.Node, to *lib.Node) (int, error) {
		cost, ok := s.EdgeCost.Get(from.ID, to.ID)
		if !ok {
			return 0, fmt.Errorf("faild to get cost. fromID=%d toID=%d", from.ID, to.ID)
		}
		return cost, nil
	}

	return search.JustCostRoutesByDijkstra(s.Graph, startName, requiredTime, getCost)
}
