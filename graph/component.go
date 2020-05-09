package graph

type ComponentProcessing interface {
	Count() int
	Connected(v0, v1 int) bool
	InComponentId(v int) int
}

type dfsComponentProcessing struct {
	g                   *UndirectedGraph
	componentCount      int
	vBelongsToComponent []int
	marked              []bool
}

func NewDfsComponentProcessing(g *UndirectedGraph) ComponentProcessing {
	dfs := &dfsComponentProcessing{
		g:                   g,
		vBelongsToComponent: make([]int, g.VertexSize()),
		marked:              make([]bool, g.VertexSize()),
	}
	dfs.init()

	return dfs
}

func (dfs *dfsComponentProcessing) init() {
	dfs.componentCount = 0
	for v := 0; v < dfs.g.VertexSize(); v++ {
		if !dfs.marked[v] {
			dfs.componentCount++
			dfs.dfs(v)
		}
	}
}

func (dfs *dfsComponentProcessing) dfs(v int) {
	dfs.marked[v] = true
	dfs.vBelongsToComponent[v] = dfs.componentCount - 1
	for _, adjacentV := range dfs.g.VerticesAdjacentTo(v) {
		if !dfs.marked[adjacentV] {
			dfs.dfs(adjacentV)
		}
	}
}

func (dfs *dfsComponentProcessing) Count() int {
	return dfs.componentCount
}

func (dfs *dfsComponentProcessing) Connected(v0, v1 int) bool {
	return dfs.vBelongsToComponent[v0] == dfs.vBelongsToComponent[v1]
}

func (dfs *dfsComponentProcessing) InComponentId(v int) int {
	return dfs.vBelongsToComponent[v]
}
