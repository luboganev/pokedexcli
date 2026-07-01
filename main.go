package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		if !scanner.Scan() {
			break
		}
		rawInput := scanner.Text()
		words := cleanInput(rawInput)
		if len(words) > 0 {
			fmt.Printf("Your command was: %s\n", words[0])
		} else {
			fmt.Print("Your command was empty\n")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Print("Failed to scan input. Ending program.\n")
	}
}
