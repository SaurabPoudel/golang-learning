package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time travel app")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2004, time.April, 20, 2, 30, 10, 10, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
