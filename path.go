package main

import (
	"container/heap"
)

type pqItem struct {
	node       Node
	priority   int
	index      int
	dirX, dirY int
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*pqItem)
	item.index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m *Map) carvePathAStar(startX, startY, endX, endY int) {
	start := Node{startX, startY}
	goal := Node{endX, endY}
	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &pqItem{node: start, priority: 0, dirX: 0, dirY: 0})
	cameFrom := make(map[Node]Node)
	costSoFar := make(map[Node]int)
	dirMap := make(map[Node][2]int)
	cameFrom[start] = start
	costSoFar[start] = 0
	dirMap[start] = [2]int{0, 0}
	dirs := []Node{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for pq.Len() > 0 {
		currItem := heap.Pop(&pq).(*pqItem)
		current := currItem.node
		if current == goal {
			break
		}

		for _, dir := range dirs {
			next := Node{current.X + dir.X, current.Y + dir.Y}
			if next.X < 1 || next.X >= Width-1 || next.Y < 1 || next.Y >= Height-1 {
				continue
			}
			moveCost := 3
			if m.Grid[next.Y][next.X] == Floor {
				moveCost = 1
			}
			if currItem.dirX != 0 || currItem.dirY != 0 {
				if currItem.dirX != dir.X || currItem.dirY != dir.Y {
					moveCost += 15
				}
			}
			newCost := costSoFar[current] + moveCost
			if prevCost, exists := costSoFar[next]; !exists || newCost < prevCost {
				costSoFar[next] = newCost
				priority := newCost + abs(goal.X-next.X) + abs(goal.Y-next.Y)
				heap.Push(&pq, &pqItem{node: next, priority: priority, dirX: dir.X, dirY: dir.Y})
				cameFrom[next] = current
				dirMap[next] = [2]int{dir.X, dir.Y}
			}
		}
	}
	curr := goal
	for curr != start {
		m.Grid[curr.Y][curr.X] = Floor
		curr = cameFrom[curr]
	}
}
