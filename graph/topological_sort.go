package graph

import (
	"errors"
)

type dependency struct {
	before string
	after  string
}

func topologicalSortProjects(dependencies []dependency) ([]string, error) {
	projectToVertex := make(map[string]int)
	vertexToProject := make([]string, 0)
	for _, d := range dependencies {
		if _, exists := projectToVertex[d.before]; !exists {
			vertexToProject = append(vertexToProject, d.before)
			projectToVertex[d.before] = len(vertexToProject) - 1
		}

		if _, exists := projectToVertex[d.after]; !exists {
			vertexToProject = append(vertexToProject, d.after)
			projectToVertex[d.after] = len(vertexToProject) - 1
		}
	}

	var paths []path
	for _, p := range dependencies {
		paths = append(paths, path{from: projectToVertex[p.before], to: projectToVertex[p.after]})
	}

	sortResult, err := topologicalSort(len(vertexToProject), paths)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, v := range sortResult {
		result = append(result, vertexToProject[v])
	}

	return result, nil
}

func topologicalSort(numberOfVertices int, paths []path) ([]int, error) {
	graph := &directedGraph{
		allEdges:     make([][]int, numberOfVertices),
		removedEdges: make([]bool, numberOfVertices),
	}
	for _, p := range paths {
		graph.add(p.from, p.to)
	}

	var result []int
	_topologicalSort(graph, &result)
	var err error
	if len(result) != numberOfVertices {
		err = errors.New("loop exists, cannot perform full sort")
	}
	return result, err
}

func _topologicalSort(graph *directedGraph, result *[]int) {
	nonAccessibleVertices := graph.getNonAccessibleVertices()
	if len(nonAccessibleVertices) == 0 {
		return
	}

	for _, v := range nonAccessibleVertices {
		graph.removeNonAccessibleVertex(v)
	}

	*result = append(*result, nonAccessibleVertices...)

	_topologicalSort(graph, result)
}

type path struct {
	from int
	to   int
}

type directedGraph struct {
	allEdges     [][]int
	removedEdges []bool
}

func (dg *directedGraph) removeNonAccessibleVertex(v int) {
	dg.allEdges[v] = nil
	dg.removedEdges[v] = true
}

func (dg *directedGraph) removeEdge(from, to int) {
	edges := dg.allEdges[from]
	for i, s := range edges {
		if s == to {
			edges[i] = edges[len(edges)-1]
			dg.allEdges[from] = edges[0 : len(edges)-1]
			return
		}
	}
}

func (dg *directedGraph) add(from, to int) {
	edges := dg.allEdges[from]
	for _, s := range edges {
		if s == to {
			return
		}
	}

	dg.allEdges[from] = append(edges, to)
}

func (dg *directedGraph) getAdjacentVertices(from int) []int {
	return dg.allEdges[from]
}

func (dg *directedGraph) getNonAccessibleVertices() []int {
	vertexAccessibility := make([]bool, len(dg.allEdges))
	for _, edges := range dg.allEdges {
		for _, v := range edges {
			vertexAccessibility[v] = true
		}
	}

	var result []int
	for vertex, accessible := range vertexAccessibility {
		if !accessible && !dg.removedEdges[vertex] {
			result = append(result, vertex)
		}
	}

	return result
}
