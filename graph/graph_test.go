package graph

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var (
	commonGraph = NewUndirectedGraph(13, [][]int{
		{0, 5},
		{4, 3},
		{0, 1},
		{9, 12},
		{6, 4},
		{5, 4},
		{0, 2},
		{11, 12},
		{9, 10},
		{0, 6},
		{7, 8},
		{9, 11},
		{5, 3},
	})
)

func TestDepthFirstSearch(t *testing.T) {
	testSearch(t, NewDepthFirstSearch)
}

func TestBreadthFirstSearch(t *testing.T) {
	testSearch(t, NewBreadthFirstSearch)
}

func testSearch(t *testing.T, newSearch func(*UndirectedGraph, int) Search) {
	rq := require.New(t)

	assertSearch(rq, newSearch(commonGraph, 0), commonGraph, map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true})
	assertSearch(rq, newSearch(commonGraph, 5), commonGraph, map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true})
	assertSearch(rq, newSearch(commonGraph, 7), commonGraph, map[int]bool{7: true, 8: true})
	assertSearch(rq, newSearch(commonGraph, 8), commonGraph, map[int]bool{7: true, 8: true})
	assertSearch(rq, newSearch(commonGraph, 9), commonGraph, map[int]bool{9: true, 10: true, 11: true, 12: true})
	assertSearch(rq, newSearch(commonGraph, 12), commonGraph, map[int]bool{9: true, 10: true, 11: true, 12: true})
}

func assertSearch(rq *require.Assertions, s Search, g *UndirectedGraph, expected map[int]bool) {
	for i := 0; i < g.VertexSize(); i++ {
		canReach, exists := expected[i]
		if !exists {
			canReach = false
		}
		rq.Equal(canReach, s.CanReach(i))

		if canReach {
			log.Printf("%d %v %d", s.Root(), s.PathTo(i), i)
		}
	}
	rq.Equal(len(expected), s.ReachableVertexCount())
}

func TestComponent(t *testing.T) {
	rq := require.New(t)
	cp := NewDfsComponentProcessing(commonGraph)
	rq.Equal(3, cp.Count())
	rq.True(cp.Connected(0, 1))
	rq.True(cp.Connected(1, 6))
	rq.False(cp.Connected(7, 9))
	rq.True(cp.Connected(10, 12))
	for v := 0; v < commonGraph.VertexSize(); v++ {
		log.Printf("%d in component %d", v, cp.InComponentId(v))
	}
}

func TestTopologicalSort(t *testing.T) {
	rq := require.New(t)
	result, err := topologicalSortProjects([]dependency{
		{"a", "d"},
		{"f", "b"},
		{"b", "d"},
		{"f", "a"},
		{"d", "c"},
	})
	rq.NoError(err)
	rq.Equal([]string{"f", "a", "b", "d", "c"}, result)

	result, err = topologicalSortProjects([]dependency{
		{"f", "c"},
		{"f", "b"},
		{"c", "a"},
		{"b", "a"},
		{"b", "e"},
		{"a", "e"},
		{"d", "g"},
	})
	rq.NoError(err)
	rq.Equal([]string{"f", "d", "c", "b", "g", "a", "e"}, result)

	result, err = topologicalSortProjects([]dependency{
		{"f", "c"},
		{"f", "b"},
		{"c", "a"},
		{"b", "a"},
		{"b", "e"},
		{"a", "e"},
		{"d", "g"},
		{"e", "f"},
	})
	rq.Error(err)
}
