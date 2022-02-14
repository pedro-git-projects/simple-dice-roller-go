package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/pedro-git-projects/dice-roller/pkg/decision"
	"github.com/pedro-git-projects/dice-roller/pkg/out"
	"github.com/pedro-git-projects/dice-roller/pkg/regular"
)

var flag bool = true

func main() {
	for flag {
		// Greeting
		out.Greet()

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		// Tests for valid options
		regular.TestInvalidInput(input)

		// Calls correct function given the input
		decision.DecideDice(input)
		if input == "q" {
			break
		}

	}

}
