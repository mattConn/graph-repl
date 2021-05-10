package commands

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type Graph simple.UndirectedGraph

var Commands = map[string]func(*simple.UndirectedGraph, []int){
	// List nodes and edges
	"ls": func(g *simple.UndirectedGraph, nodes []int) {
		fmt.Println(graph.NodesOf(g.Nodes()))
		fmt.Println(graph.EdgesOf(g.Edges()))
	},

	// Clear nodes
	"cln": func(g *simple.UndirectedGraph, nodes []int) {
		for _, n := range graph.NodesOf(g.Nodes()) {
			g.RemoveNode(n.ID())
		}
	},

	// Clear edges
	"cle": func(g *simple.UndirectedGraph, nodes []int) {
		for _, e := range graph.EdgesOf(g.Edges()) {
			g.RemoveEdge(e.From().ID(), e.To().ID())
		}
	},

	// Add nodes over range
	"..n": func(g *simple.UndirectedGraph, nodes []int) {
		from := nodes[0]
		to := nodes[1]
		for ; from <= to; from++ {
			g.AddNode(simple.Node(int64(from)))
		}
	},

	// Add nodes
	"addn": func(g *simple.UndirectedGraph, nodes []int) {
		for _, n := range nodes {
			g.AddNode(simple.Node(int64(n)))
		}
	},

	// Delete nodes
	"deln": func(g *simple.UndirectedGraph, nodes []int) {
		for _, n := range nodes {
			g.RemoveNode(int64(n))
		}
	},

	// Add edges
	"adde": func(g *simple.UndirectedGraph, nodes []int) {
		for i := 0; i < len(nodes)-1; i += 2 {
			g.SetEdge(g.NewEdge(simple.Node(nodes[i]), simple.Node(nodes[i+1])))
		}
	},

	// Delete edges
	"dele": func(g *simple.UndirectedGraph, nodes []int) {
		for i := 0; i < len(nodes)-1; i += 2 {
			g.RemoveEdge(int64(nodes[i]), int64(nodes[i+1]))
		}
	},
}
