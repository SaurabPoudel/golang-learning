package main

import "fmt"

func main() {

	fmt.Println("Welcome to array in golangs")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Grapes"

	fmt.Println("Fruit list is: ", fruitList)

	fmt.Println("Fruit List is:", len(fruitList))

	var vegList = [3]string{"vindi", "mushroom", "pyaj"}

	fmt.Println(vegList)

}
