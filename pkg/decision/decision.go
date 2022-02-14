package decision

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pedro-git-projects/dice-roller/pkg/dice"
)

func DecideDice(input string) {
	if input == "a" {
		d4 := dice.Dice{dice.D4}
		fmt.Println("You rolled: ", d4.Roll())
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	if input == "b" {
		d6 := dice.Dice{dice.D6}
		fmt.Println("You rolled: ", d6.Roll())
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	if input == "c" {
		d8 := dice.Dice{dice.D8}
		fmt.Println("You rolled: ", d8.Roll())
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	if input == "d" {
		d20 := dice.Dice{dice.D20}
		fmt.Println("You rolled: ", d20.Roll())
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	if input == "e" {
		d100 := dice.Dice{dice.D100}
		fmt.Println("You rolled: ", d100.Roll())
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
