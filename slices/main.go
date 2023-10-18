package main

import "fmt"

func main() {
	//var admin []string
	////fmt.Println(len(admin))
	////fmt.Println(cap(admin))
	////fmt.Println(admin)
	//admin = append(admin, "Saurab", "Santosh", "Ranjit", "Shishir", "Frecklehead")
	////fmt.Println("Admin of this website  are ", admin)
	//
	//super_admin := admin[0:3]
	////fmt.Println("Super admins are: ", super_admin)
	////fmt.Printf("length = %d\n", len(super_admin))
	////fmt.Printf("capacity = %d\n", cap(super_admin))
	//
	////myslice1 := make([]int, 5, 10)
	////fmt.Printf("myslice1 = %v\n", myslice1)
	////fmt.Printf("length = %d\n", len(myslice1))
	////fmt.Printf("capacity = %d\n", cap(myslice1))
	////
	////// with omitted capacity
	////myslice2 := make([]int, 5)
	////fmt.Printf("myslice2 = %v\n", myslice2)
	////fmt.Printf("length = %d\n", len(myslice2))
	////fmt.Printf("capacity = %d\n", cap(myslice2))
	//
	//users := append(admin, super_admin...)
	//fmt.Printf(" users= %v\n", users)
	//fmt.Printf("length = %d\n", len(users))
	//fmt.Printf("capacity = %d\n", cap(users))

	//Memory Efficiency
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
