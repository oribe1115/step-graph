package lib

// EdgeCost 辺のコスト情報を保持する
type EdgeCost struct {
	Costs      map[int]map[int]int
	IsDirected bool // 有向グラフかどうか
}

// InitEdgeCost 初期化したEdgeCostを返す
func InitEdgeCost(isDirected bool) *EdgeCost {
	return &EdgeCost{
		Costs:      map[int]map[int]int{},
		IsDirected: isDirected,
	}
}

// Set EdgeCostに辺のコスト情報を記録する
func (e *EdgeCost) Set(firstID int, secondID int, cost int) {
	firstID, secondID = e.checkIDOrder(firstID, secondID)

	secondMap, ok := e.Costs[firstID]
	if !ok {
		e.Costs[firstID] = map[int]int{}
		secondMap = e.Costs[firstID]
	}

	secondMap[secondID] = cost
}

// Get EdgeCostからコスト情報を取り出す
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
