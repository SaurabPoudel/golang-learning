package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// syntax for input
	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our product: ")
	// comma ok syntax || err ok

	// if you care about input and dc about error then input, _ is used
	// if you dc about input but care about error then _, input is used
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
}
