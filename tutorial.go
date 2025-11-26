package main

import "fmt"

func main() {
	// printing
	fmt.Println("Starting Textio server...")

	// declaring variables, which sets them to their 0 value by default
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	// formatted print statements
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// function call, which doesn't need to be declared beforehand apparently
	birthdayGreeting()

	// type casting works as expected
	moneyFloat := 2.34
	moneyInt := int(moneyFloat)

	fmt.Println(moneyInt, "is the int of", moneyFloat)
}

func birthdayGreeting() {
	/*
		better way of declaring variables and assigning immediately
		this way can't be used outside functions, though
	*/
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)
}
