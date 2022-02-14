package regular

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func matchRunes(input string) bool {
	reg, err := regexp.Compile("[A-Ea-e]")
	if err != nil {
		log.Fatal(err)
	}
	return reg.MatchString(input)
}

func TestInvalidInput(input string) {
	if matchRunes(input) == false && input != "q" {
		fmt.Println("Please enter a valid option")
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
