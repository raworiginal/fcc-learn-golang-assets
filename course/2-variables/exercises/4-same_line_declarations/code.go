package main

import "fmt"

func main() {
	// declare here
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Printf("the type of averagerOpenRate is %T", averageOpenRate)
	fmt.Println()
	fmt.Println(averageOpenRate, displayMessage)
}
