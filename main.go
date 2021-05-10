package main

import (
	"bufio"
	"fmt"
	"graph-repl/commands"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var (
		input    string
		tokens   []string
		operands []int
		cmd      func(*simple.UndirectedGraph, []int)
		fpair    commands.Fpair
		ok       bool
	)
	g := simple.NewUndirectedGraph()

	for true {
		fmt.Printf("> ")

		// Read stdin and split into tokens
		input, _ = reader.ReadString('\n')
		tokens = strings.Fields(input)
		if len(tokens) == 0 {
			continue
		}
		if tokens[0] == "quit" {
			return
		}
		if tokens[0] == "help" {
			for key, pair := range commands.Commands {
				fmt.Printf("%s - %s\n", key, pair.Desc)
			}
			continue
		}

		// Check if command exists
		if fpair, ok = commands.Commands[tokens[0]]; !ok {
			fmt.Println("Unrecognized command:", tokens[0])
			continue
		}
		cmd = fpair.Fn

		// Convert tokens to ints
		operands = make([]int, len(tokens[1:]))
		i := 0
		for _, tok := range tokens[1:] {
			n, err := strconv.Atoi(tok)

			if err != nil {
				fmt.Println("Invalid operand:", tok)
				break
			}

			operands[i] = n
			i++
		}

		// Exec command
		cmd(g, operands)
		if tokens[0] != "ls" {
			commands.Render(g)
		}
	}
}
