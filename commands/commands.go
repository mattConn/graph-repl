package commands

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
)

var Commands = map[string]func(graph.Graph, []int){
	"test": func(graph.Graph, []int) {
		fmt.Println("TEST FN")
	},
}
