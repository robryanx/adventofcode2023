package util

import (
	hp "container/heap"
	"strconv"
	"strings"
)

type path struct {
	value int
	nodes []string
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Heap struct {
	values *minPath
}

func newHeap() *Heap {
	return &Heap{values: &minPath{}}
}

func (h *Heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *Heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}

type Edge struct {
	node   string
	weight int
}

type Graph struct {
	nodes map[string][]Edge
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]Edge)}
}

func (g *Graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], Edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], Edge{node: origin, weight: weight})
}

func (g *Graph) getEdges(node string) []Edge {
	return g.nodes[node]
}

func (g *Graph) GetPath(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				newNodes := append([]string{}, append(p.nodes, e.node)...)

				if validPath(newNodes) {
					h.push(path{value: p.value + e.weight, nodes: newNodes})
				}
			}
		}

		visited[node] = true
	}

	return 0, nil
}

func validPath(nodes []string) bool {
	last := nodes[len(nodes)-1]
	y, x := coordFromString(last)
	dir := Unknown
	count := 0

	for i := len(nodes) - 2; i >= 0; i-- {
		found := false
		nodeY, nodeX := coordFromString(nodes[i])
		if (dir == Unknown || dir == North) && nodeX == x && nodeY < y {
			dir = North
			count += y - nodeY
			found = true
		}
		if (dir == Unknown || dir == South) && nodeX == x && nodeY > y {
			dir = South
			count += nodeY - y
			found = true
		}
		if (dir == Unknown || dir == West) && nodeX > x && nodeY == y {
			dir = West
			count += nodeX - x
			found = true
		}
		if (dir == Unknown || dir == East) && nodeX < x && nodeY == y {
			dir = East
			count += x - nodeX
			found = true
		}

		if !found {
			break
		}
	}

	return count < 3
}

func coordFromString(s string) (int, int) {
	parts := strings.Split(s, "-")
	y, _ := strconv.Atoi(parts[0])
	x, _ := strconv.Atoi(parts[1])

	return y, x
}
