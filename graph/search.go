package graph

import "github.com/martincui/algorithm/datastructure"

type Search interface {
	CanReach(targetVertex int) bool
	ReachableVertexCount() int
	PathTo(targetVertex int) []int
	Root() int
}

type depthFirstSearch struct {
	g               *UndirectedGraph
	root            int
	canReachTo      []bool
	pathTo          []int
	canReachToCount int
}

func NewDepthFirstSearch(g *UndirectedGraph, root int) Search {
	dfs := &depthFirstSearch{
		g:          g,
		root:       root,
		canReachTo: make([]bool, g.VertexSize()),
		pathTo:     make([]int, g.VertexSize()),
	}
	dfs.dfs(root)
	return dfs
}

func (dfs *depthFirstSearch) dfs(v int) {
	dfs.canReachTo[v] = true
	dfs.canReachToCount++
	for _, adjacentV := range dfs.g.VerticesAdjacentTo(v) {
		if !dfs.canReachTo[adjacentV] {
			dfs.pathTo[adjacentV] = v
			dfs.dfs(adjacentV)
		}
	}
}

func (dfs *depthFirstSearch) Root() int{
	return dfs.root
}

func (dfs *depthFirstSearch) CanReach(targetVertex int) bool {
	return dfs.canReachTo[targetVertex]
}

func (dfs *depthFirstSearch) PathTo(targetVertex int) []int {
	if !dfs.canReachTo[targetVertex] {
		return nil
	}

	stack := datastructure.NewStack()
	for v := targetVertex; v != dfs.root; v = dfs.pathTo[v] {
		stack.Push(v)
	}
	var path []int
	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		path = append(path, v.(int))
	}
	return path
}

func (dfs *depthFirstSearch) ReachableVertexCount() int {
	return dfs.canReachToCount
}

type breadthFirstSearch struct {
	g               *UndirectedGraph
	root            int
	canReachTo      []bool
	pathTo          []int
	canReachToCount int
}

func NewBreadthFirstSearch(g *UndirectedGraph, root int) Search {
	bfs := &breadthFirstSearch{
		g:          g,
		root:       root,
		canReachTo: make([]bool, g.VertexSize()),
		pathTo:     make([]int, g.VertexSize()),
	}
	bfs.bfs()
	return bfs
}

func (bfs *breadthFirstSearch) bfs() {
	toProcessVertices := datastructure.NewQueue()
	toProcessVertices.Add(bfs.root)

	for toProcessV, err := toProcessVertices.Pop(); err == nil; toProcessV, err = toProcessVertices.Pop() {
		bfs.canReachToCount++
		bfs.canReachTo[toProcessV.(int)] = true
		for _, adjacentV := range bfs.g.VerticesAdjacentTo(toProcessV.(int)) {
			if !bfs.canReachTo[adjacentV] {
				toProcessVertices.Add(adjacentV)
				bfs.canReachTo[adjacentV] = true
				bfs.pathTo[adjacentV] = toProcessV.(int)
			}
		}
	}
}

func (bfs *breadthFirstSearch) Root() int{
	return bfs.root
}

func (bfs *breadthFirstSearch) CanReach(targetVertex int) bool {
	return bfs.canReachTo[targetVertex]
}

func (bfs *breadthFirstSearch) ReachableVertexCount() int {
	return bfs.canReachToCount
}

func (bfs *breadthFirstSearch) PathTo(targetVertex int) []int {
	if !bfs.canReachTo[targetVertex] {
		return nil
	}

	stack := datastructure.NewStack()
	for v := targetVertex; v != bfs.root; v = bfs.pathTo[v] {
		stack.Push(v)
	}
	var path []int
	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		path = append(path, v.(int))
	}
	return path
}
