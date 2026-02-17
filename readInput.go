package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	fmt.Printf("You entered: %s", input)

	// add reading integers
	// add reading floats
	// how to get byte
	// how to get rune
}
