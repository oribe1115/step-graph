package lib

type EdgeCost struct {
	Costs      map[int]map[int]int
	IsDirected bool // 有向グラフかどうか
}

func InitEdgeCost(isDirected bool) *EdgeCost {
	return &EdgeCost{
		Costs:      map[int]map[int]int{},
		IsDirected: isDirected,
	}
}

func (e *EdgeCost) Set(firstID int, secondID int, cost int) {
	firstID, secondID = e.checkIDOrder(firstID, secondID)

	secondMap, ok := e.Costs[firstID]
	if !ok {
		e.Costs[firstID] = map[int]int{}
		secondMap = e.Costs[firstID]
	}

	secondMap[secondID] = cost
}

func (e *EdgeCost) Get(firstID int, secondID int) (cost int, ok bool) {
	firstID, secondID = e.checkIDOrder(firstID, secondID)
	secondMap, ok := e.Costs[firstID]
	if !ok {
		return 0, false
	}

	cost, ok = secondMap[secondID]
	return cost, ok
}

func (e *EdgeCost) checkIDOrder(idA int, idB int) (firstID int, secondID int) {
	if !e.IsDirected && idA > idB {
		return idB, idA
	}
	return idA, idB
}
