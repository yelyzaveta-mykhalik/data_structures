package main

import "fmt"

// representation of adjacency list graph
type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

//adding vertex to graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v did not add beacause it is an existing key")
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Edge (%v-->%v) already exists", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(l []*Vertex, k int) bool {
	for _, v := range l {
		if k == v.key {
			return true
		}
	}
	return false
}

//printing adjacency list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}
func main() {
	testGraph := &Graph{}

	for i := 0; i < 5; i++ {
		testGraph.AddVertex(i)
	}

	testGraph.AddEdge(1, 4)
	testGraph.AddEdge(7, 1)
	testGraph.AddEdge(1, 4)
	testGraph.Print()
}
