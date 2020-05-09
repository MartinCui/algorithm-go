package graph

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {
	testSearch(t, NewDepthFirstSearch)
}

func TestBreadthFirstSearch(t *testing.T) {
	testSearch(t, NewBreadthFirstSearch)
}

func testSearch(t *testing.T, newSearch func(*UndirectedGraph, int) Search) {
	rq := require.New(t)
	g := NewUndirectedGraph(13, [][]int{
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
	assertSearch(rq, newSearch(g, 0), g, map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true})
	assertSearch(rq, newSearch(g, 5), g, map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true})
	assertSearch(rq, newSearch(g, 7), g, map[int]bool{7: true, 8: true})
	assertSearch(rq, newSearch(g, 8), g, map[int]bool{7: true, 8: true})
	assertSearch(rq, newSearch(g, 9), g, map[int]bool{9: true, 10: true, 11: true, 12: true})
	assertSearch(rq, newSearch(g, 12), g, map[int]bool{9: true, 10: true, 11: true, 12: true})
}

func assertSearch(rq *require.Assertions, s Search, g *UndirectedGraph, expected map[int]bool) {
	for i := 0; i < g.VertexSize(); i++ {
		canReach, exists := expected[i]
		if !exists {
			canReach = false
		}
		rq.Equal(canReach, s.CanReach(i))

		if canReach{
			log.Printf("%d %v %d", s.Root(), s.PathTo(i), i)
		}
	}
	rq.Equal(len(expected), s.ReachableVertexCount())
}
