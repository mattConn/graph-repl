package commands

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/graphs/gen"
	"gonum.org/v1/gonum/graph/simple"
)

// Graph function paired with string description
type Fpair struct {
	Desc string
	Fn   func(*simple.UndirectedGraph, []int)
}

var Commands = map[string]Fpair{
	// List nodes and edges
	"ls": {
		Fn:   list,
		Desc: "List all nodes and edges",
	},

	// Clear nodes
	"cln": {
		Fn:   clearNodes,
		Desc: "Remove all nodes",
	},

	// Clear edges
	"cle": {
		Fn:   clearEdges,
		Desc: "Remove all edges",
	},

	// Add nodes over range
	"..": {
		Fn:   addNodesRange,
		Desc: "Add nodes over a range, e.g. .. 1 5",
	},

	// Add nodes
	"addn": {
		Fn:   addNodes,
		Desc: "Add nodes",
	},

	// Remove nodes
	"deln": {
		Fn:   removeNodes,
		Desc: "Remove nodes",
	},

	// Add edges
	"adde": {
		Fn:   addEdges,
		Desc: "Add edges",
	},

	// Delete edges
	"dele": {
		Fn:   removeEdges,
		Desc: "Remove edges",
	},

	// Add cycle
	"addc": {
		Fn:   addCycle,
		Desc: "Add cycle",
	},

	// Add star
	"adds": {
		Fn:   addStar,
		Desc: "Add star",
	},

	// Add complete
	"addk": {
		Fn:   addComplete,
		Desc: "Add complete",
	},

	// Add wheel
	"addw": {
		Fn:   addWheel,
		Desc: "Add wheel",
	},

	// Add path
	"addp": {
		Fn:   addPath,
		Desc: "Add path",
	},
}

type nodeList []int

func (n nodeList) ID(i int) int64 {
	return int64(n[i])
}

func (n nodeList) Len() int {
	return len(n)
}

func list(g *simple.UndirectedGraph, nodes []int) {
	fmt.Println(graph.NodesOf(g.Nodes()))
	fmt.Println(graph.EdgesOf(g.Edges()))
}

func clearNodes(g *simple.UndirectedGraph, nodes []int) {
	for _, n := range graph.NodesOf(g.Nodes()) {
		g.RemoveNode(n.ID())
	}
}

func clearEdges(g *simple.UndirectedGraph, nodes []int) {
	for _, e := range graph.EdgesOf(g.Edges()) {
		g.RemoveEdge(e.From().ID(), e.To().ID())
	}
}

func addNodesRange(g *simple.UndirectedGraph, nodes []int) {
	from := nodes[0]
	to := nodes[1]
	for ; from <= to; from++ {
		g.AddNode(simple.Node(int64(from)))
	}
}

func addNodes(g *simple.UndirectedGraph, nodes []int) {
	for _, n := range nodes {
		g.AddNode(simple.Node(int64(n)))
	}
}

func removeNodes(g *simple.UndirectedGraph, nodes []int) {
	for _, n := range nodes {
		g.RemoveNode(int64(n))
	}
}

func addEdges(g *simple.UndirectedGraph, nodes []int) {
	for i := 0; i < len(nodes)-1; i += 2 {
		g.SetEdge(g.NewEdge(simple.Node(nodes[i]), simple.Node(nodes[i+1])))
	}
}

func removeEdges(g *simple.UndirectedGraph, nodes []int) {
	for i := 0; i < len(nodes)-1; i += 2 {
		g.RemoveEdge(int64(nodes[i]), int64(nodes[i+1]))
	}
}

func addCycle(g *simple.UndirectedGraph, nodes []int) {
	gen.Cycle(g, nodeList(nodes))
}

func addStar(g *simple.UndirectedGraph, nodes []int) {
	gen.Star(g, int64(nodes[0]), nodeList(nodes[1:]))
}

func addComplete(g *simple.UndirectedGraph, nodes []int) {
	gen.Complete(g, nodeList(nodes))
}

func addWheel(g *simple.UndirectedGraph, nodes []int) {
	gen.Wheel(g, int64(nodes[0]), nodeList(nodes[1:]))
}

func addPath(g *simple.UndirectedGraph, nodes []int) {
	gen.Path(g, nodeList(nodes))
}
