package graph

func NewUndirectedGraph(verticesNumber int, edges [][]int) *UndirectedGraph {
	ug := &UndirectedGraph{connectedVertices: make([][]int, verticesNumber)}
	for _, edge := range edges {
		ug.AddEdge(edge[0], edge[1])
	}
	return ug
}

type UndirectedGraph struct {
	edgeSize          int
	connectedVertices [][]int
}

func (ug *UndirectedGraph) EdgeSize() int {
	return ug.edgeSize
}

func (ug *UndirectedGraph) VertexSize() int {
	return len(ug.connectedVertices)
}

func (ug *UndirectedGraph) AddEdge(v0, v1 int) {
	ug.connectedVertices[v0] = append(ug.connectedVertices[v0], v1)
	ug.connectedVertices[v1] = append(ug.connectedVertices[v1], v0)
	ug.edgeSize++
}

func (ug *UndirectedGraph) VerticesAdjacentTo(v int) []int {
	return ug.connectedVertices[v]
}
