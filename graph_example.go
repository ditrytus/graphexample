package main

import (
	"fmt"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

type MyDomainType struct {
	id int64
}

func (m MyDomainType) ID() int64 {
	return m.id
}

func main() {
	a := MyDomainType{id: 1}
	b := MyDomainType{id: 2}
	c := MyDomainType{id: 3}

	g := simple.NewUndirectedGraph()
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)

	g.SetEdge(g.NewEdge(a, b))
	g.SetEdge(g.NewEdge(a, c))

	shortest, _ := path.AStar(b, c, g, path.NullHeuristic)
	path, _ := shortest.To(c.ID())
	fmt.Println(path)
}
