// http://play.golang.org/p/W5qjLAsJbI

// Sample program to show the basic concept of pass by value.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("Before:", count, &count)

	// Pass the "value of" the variable count.
	inc(count)

	fmt.Println("After:", count, &count)
}

// inc declares count as a variable whose value is
// always an integer.
func inc(count int) {
	// Increment the "value of" count.
	count++
	fmt.Println("Inc:", count, &count)
}
