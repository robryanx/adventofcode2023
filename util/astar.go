package util

import (
	"container/heap"
	"fmt"
)

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].rank < pq[j].rank
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	no := x.(*node)
	no.index = n
	*pq = append(*pq, no)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	no := old[n-1]
	no.index = -1
	*pq = old[0 : n-1]
	return no
}

type NodePos struct {
	Y         int
	X         int
	direction Direction
	cost      int
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

// node is a wrapper to store A* data for a Pather node.
type node struct {
	pos    NodePos
	rank   int
	parent *node
	open   bool
	closed bool
	index  int
}

func (p node) parentDirection() (int, Direction) {
	if p.parent == nil {
		return 0, Unknown
	}

	if p.pos.X == p.parent.pos.X && p.pos.Y < p.parent.pos.Y {
		count := 1
		current := p.parent
		for current.parent != nil {
			if current.pos.X == current.parent.pos.X && current.pos.Y < current.parent.pos.Y {
				count++
			} else {
				break
			}
			current = current.parent
		}

		return count, North
	}

	if p.pos.X == p.parent.pos.X && p.pos.Y > p.parent.pos.Y {
		count := 1
		current := p.parent
		for current.parent != nil {
			if current.pos.X == current.parent.pos.X && current.pos.Y > current.parent.pos.Y {
				count++
			} else {
				break
			}
			current = current.parent
		}

		return count, South
	}

	if p.pos.Y == p.parent.pos.Y && p.pos.X > p.parent.pos.X {
		count := 1
		current := p.parent
		for current.parent != nil {
			if current.pos.Y == current.parent.pos.Y && current.pos.X > current.parent.pos.X {
				count++
			} else {
				break
			}
			current = current.parent
		}

		return count, East
	}

	if p.pos.Y == p.parent.pos.Y && p.pos.X < p.parent.pos.X {
		count := 1
		current := p.parent
		for current.parent != nil {
			if current.pos.Y == current.parent.pos.Y && current.pos.X < current.parent.pos.X {
				count++
			} else {
				break
			}
			current = current.parent
		}

		return count, West
	}

	return 0, Unknown
}

func (p node) neighborsWithInterval(grid [][]byte) []NodePos {
	list := make([]NodePos, 0, 12)

	_, parentDirection := p.parentDirection()
	if parentDirection != North {
		for i := 1; i < 4; i++ {
			if p.pos.Y-i >= 0 {
				newCost := 0
				for j := 1; j <= i; j++ {
					newCost += int(grid[p.pos.Y-j][p.pos.X] - '0')
				}

				list = append(list, NodePos{
					Y:    p.pos.Y - i,
					X:    p.pos.X,
					cost: newCost,
				})
			}
		}
	}

	if parentDirection != South {
		for i := 1; i < 4; i++ {
			if p.pos.Y+i < len(grid) {
				newCost := 0
				for j := 1; j <= i; j++ {
					newCost += int(grid[p.pos.Y+j][p.pos.X] - '0')
				}

				list = append(list, NodePos{
					Y:    p.pos.Y + i,
					X:    p.pos.X,
					cost: newCost,
				})
			}
		}
	}

	if parentDirection != West {
		for i := 1; i < 4; i++ {
			if p.pos.X-i >= 0 {
				newCost := 0
				for j := 1; j <= i; j++ {
					newCost += int(grid[p.pos.Y][p.pos.X-j] - '0')
				}

				list = append(list, NodePos{
					Y:    p.pos.Y,
					X:    p.pos.X - i,
					cost: newCost,
				})
			}
		}
	}

	if parentDirection != East {
		for i := 1; i < 4; i++ {
			if p.pos.X+i < len(grid[0]) {
				newCost := 0
				for j := 1; j <= i; j++ {
					newCost += int(grid[p.pos.Y][p.pos.X+j] - '0')
				}

				list = append(list, NodePos{
					Y:    p.pos.Y,
					X:    p.pos.X + i,
					cost: newCost,
				})
			}
		}
	}

	return list
}

func (p node) neighbors(grid [][]byte) []NodePos {
	list := make([]NodePos, 0, 4)
	maxCount := 3

	count, parentDirection := p.parentDirection()
	if parentDirection != South && (parentDirection != North || count < maxCount) {
		if p.pos.Y > 0 {
			list = append(list, NodePos{
				Y:    p.pos.Y - 1,
				X:    p.pos.X,
				cost: int(grid[p.pos.Y-1][p.pos.X] - '0'),
			})
		}
	}

	if parentDirection != North && (parentDirection != South || count < maxCount) {
		if p.pos.Y < len(grid)-1 {
			list = append(list, NodePos{
				Y:    p.pos.Y + 1,
				X:    p.pos.X,
				cost: int(grid[p.pos.Y+1][p.pos.X] - '0'),
			})
		}
	}

	if parentDirection != East && (parentDirection != West || count < maxCount) {
		if p.pos.X > 0 {
			list = append(list, NodePos{
				Y:    p.pos.Y,
				X:    p.pos.X - 1,
				cost: int(grid[p.pos.Y][p.pos.X-1] - '0'),
			})
		}
	}

	if parentDirection != West && (parentDirection != East || count < maxCount) {
		if p.pos.X < len(grid[0])-1 {
			list = append(list, NodePos{
				Y:    p.pos.Y,
				X:    p.pos.X + 1,
				cost: int(grid[p.pos.Y][p.pos.X+1] - '0'),
			})
		}
	}

	return list
}

func (p node) estimatedCost(grid [][]byte, to NodePos) int {
	return abs(p.pos.Y-to.Y) + abs(p.pos.X-to.X)
}

type nodeMap map[int]*node

func (nm nodeMap) get(pos NodePos) *node {
	n, ok := nm[pos.Y*1000+pos.X]
	if !ok {
		n = &node{
			pos: pos,
		}
		nm[pos.Y*1000+pos.X] = n
	}
	return n
}

func Pathfind(grid [][]byte, from NodePos, to NodePos) (path []NodePos, distance int, found bool) {
	nm := nodeMap{}
	nq := &priorityQueue{}
	heap.Init(nq)

	fromNode := nm.get(from)
	fromNode.open = true
	heap.Push(nq, fromNode)
	for {
		if nq.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		if current.pos.Y == 0 && current.pos.X == 2 {
			fmt.Printf("%+v\n", current)
			fmt.Println(current.parentDirection())
			fmt.Println(current.neighbors(grid))
		}

		if current == nm.get(to) {
			// Found a path to the goal.
			p := []NodePos{}
			curr := current
			for curr != nil {
				_, dir := curr.parentDirection()
				curr.pos.direction = dir
				p = append(p, curr.pos)
				curr = curr.parent
			}
			return p, current.pos.cost, true
		}

		for _, neighbor := range current.neighbors(grid) {
			cost := current.pos.cost + neighbor.cost
			neighborNode := nm.get(neighbor)

			if current.pos.Y == 0 && current.pos.X == 2 {
				fmt.Printf("%+v\n", neighborNode)
				fmt.Println(cost)
			}

			if cost <= neighborNode.pos.cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.pos.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost
				neighborNode.parent = current

				heap.Push(nq, neighborNode)
			}
		}
	}
}

type Direction int

const (
	Unknown Direction = iota
	North
	East
	South
	West
)
