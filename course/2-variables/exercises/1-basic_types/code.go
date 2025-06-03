package main

import "fmt"

var (
	smsSendingLimit int
	costPerSMS      float64
	hasPermission   bool
	username        string
)

func main() {
	// initialize variables here

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
