package main

import "fmt"

// Not allowed
// jwtToken := 30000

// LoginToken public variable first letter capital
const LoginToken string = "sdfakld"

func main() {
	var username string = "Saurab"
	fmt.Println(username)
	fmt.Printf("Variable if of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable if of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable if of type : %T \n", smallVal)

	var smallFloat float32 = 255.455555
	fmt.Println(smallFloat)
	fmt.Printf("Variable if of type : %T \n", smallFloat)

	// default values and some aliases
	var anotherVaribable int
	fmt.Println(anotherVaribable)
	fmt.Printf("Variable is of type: %T \n", anotherVaribable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type :%T \n", anotherString)

	// implicit type

	var website = "poudelsaurab20.com.np"
	fmt.Println(website)

	// no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
