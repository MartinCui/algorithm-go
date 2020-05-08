package graph

import "github.com/martincui/algorithm/datastructure"

type Search interface {
	CanReach(targetVertex int) bool
	ReachableVertexCount() int
}

type depthFirstSearch struct {
	g               *UndirectedGraph
	root            int
	canReachTo      []bool
	canReachToCount int
}

func NewDepthFirstSearch(g *UndirectedGraph, root int) Search {
	canReachTo := make([]bool, g.VertexSize())
	dfs := &depthFirstSearch{g: g, root: root, canReachTo: canReachTo}
	dfs.dfs(root)
	return dfs
}

func (dfs *depthFirstSearch) dfs(targetVertex int) {
	dfs.canReachTo[targetVertex] = true
	dfs.canReachToCount++
	for _, adjacentV := range dfs.g.VerticesAdjacentTo(targetVertex) {
		if !dfs.canReachTo[adjacentV] {
			dfs.dfs(adjacentV)
		}
	}
}

func (dfs *depthFirstSearch) CanReach(targetVertex int) bool {
	return dfs.canReachTo[targetVertex]
}

func (dfs *depthFirstSearch) ReachableVertexCount() int {
	return dfs.canReachToCount
}

type breadthFirstSearch struct {
	g               *UndirectedGraph
	root            int
	canReachTo      []bool
	canReachToCount int
}

func NewBreadthFirstSearch(g *UndirectedGraph, root int) Search {
	canReachTo := make([]bool, g.VertexSize())
	bfs := &breadthFirstSearch{g: g, root: root, canReachTo: canReachTo}
	bfs.bfs()
	return bfs
}

func (bfs *breadthFirstSearch) bfs() {
	toProcessVertices := datastructure.NewQueue()
	toProcessVertices.Add(bfs.root)

	for toProcessV, err := toProcessVertices.Pop(); err == nil; toProcessV, err = toProcessVertices.Pop() {
		if bfs.canReachTo[toProcessV.(int)] {
			continue
		}

		bfs.canReachToCount++
		bfs.canReachTo[toProcessV.(int)] = true
		for _, adjacentV := range bfs.g.VerticesAdjacentTo(toProcessV.(int)) {
			if !bfs.canReachTo[adjacentV] {
				toProcessVertices.Add(adjacentV)
			}
		}
	}
}

func (bfs *breadthFirstSearch) CanReach(targetVertex int) bool {
	return bfs.canReachTo[targetVertex]
}

func (bfs *breadthFirstSearch) ReachableVertexCount() int {
	return bfs.canReachToCount
}
