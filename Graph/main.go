//Adjacent list

package main

import "fmt"

//Graph struct
type Graph struct {
	vertices []*Vertex
}

//Vertex strct
type Vertex struct {
	key      int
	adjacent []*Vertex
}

//AddVertex adds vertex to the graph
func (g *Graph) AddVertex(k int) {
	if Contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

//AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check the errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid Edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if Contains(fromVertex.adjacent, to) {
		//check if already exist or not
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

//getVertex returns a pointer to the vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//Contains to check out the key is added before or not
func Contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

//Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}

	for i := 0; i <= 5; i++ {
		test.AddVertex(i)
	}

	// fmt.Println(test)

	// test.AddVertex(0)
	// test.AddVertex(1)
	// test.AddVertex(2)
	//to check if existing key can be added or not

	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.AddEdge(6, 2)
	test.AddEdge(3, 4)
	test.AddEdge(5, 1)
	test.AddEdge(2, 5)

	test.Print()

}
