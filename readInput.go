package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter something: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered: ", input)

	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error reading number: ", err)
		return
	}
	fmt.Println("You entered: ", num)
	multiplied := 2 * num
	fmt.Println(multiplied)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press a key: ")

	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	fmt.Printf("You pressed: %c\n", char)

}
