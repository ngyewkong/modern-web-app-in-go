// package declaration by convention is main
package main

// import the fmt package from the std lib
import (
	"fmt"
	"log"
)

// package level variables declared outside out of the main function
var myName string

// necessary main function
// can only have one main() function and the main function takes no args & return nth
// to run this code (terminal execute go run fileName.go)
func main() {
	// print string to console
	fmt.Println("Hello, world.")

	// variables declaration
	// variable signature
	// var variableName dataType
	// unused declared variables will have warning by the go compiler
	var whatToSay string = "Bonjour"

	fmt.Println(whatToSay)

	// int data type -> int by default is int64 for base64 arch computers
	var someNum int

	someNum = 5

	// PrintLn do not need to leave a blank space (by default the var will have a blank space in front of the initial string)
	// this is printed -> The number is 5 instead of The number is5
	fmt.Println("The number is", someNum)

	// call the function that is declared outside of the main function
	saySomething()

	// variable declaration shorthand
	// what is returned from the function will be assigned to the variable that is newly declared with := notation
	whatWasSaid := saySomething()

	// this print The function returned something to the console
	fmt.Println("The function returned", whatWasSaid)

	repeatedLines, numOfTimes := sayMultipleThings()

	// print "This is repeated 8"
	fmt.Println(repeatedLines, numOfTimes)

	// showing use of pointers
	var myString string
	myString = "GREEN"

	log.Println("Colour is now set to", myString)
	// call the pointer function
	// pointer is just smth that points to a specific location in memory and gets that particular location in memory
	// so instead of passing the parameter myString as an actual variable
	// we have to pass a reference to that variable
	// using & ampersand (&myString) -> reference
	// * -> pointer
	changeUsingPointer(&myString)

	// myString actually become RED despite the function not returning anything
	log.Println("After func call, colour is now set to", myString)
}

// other functions can be declared outside of the main function
// follow signature
// func someFunction() returnDataType {
// return ""
// }

func saySomething() string {
	return "something"
}

// function can return multiple objects
func sayMultipleThings() (string, int) {
	return "This is repeated", 8
}

// Pointers in golang
func changeUsingPointer(s *string) {

	// this s will be the memory address when the reference of the variable is being passed in
	log.Println("s pointer is set to", s)
	newValue := "RED"

	// this means go to that memory address and change the contents from whatever it used to be to what I set in newValue (RED)
	*s = newValue
}
