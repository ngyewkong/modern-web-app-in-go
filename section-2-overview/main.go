// package declaration by convention is main
package main

// import the fmt package from the std lib
import (
	"fmt"
	"log"
	"time"
)

// package level variables declared outside out of the main function
var myName string

// other ways to declare/init variables
var numString = "seven" // go compiler auto infer numString to be string

// Custom Data Type using Struct
// when primitive data types are not enough
// eg injecting a Person into a db (firstName, lastName, phoneNum, age, birthDate)
// signature
// type nameOfStruct struct {}
type User struct {
	// go do not have OOP (Public, Private, Protected)
	// so if CapFirstLetter -> accessible/visible outside of this package
	// then if smallFirstLetter -> within the package is accessible
	// eg. log.Fatalf is visible outside of the log package
	FirstName string
	LastName  string
	PhoneNum  string
	Age       int
	BirthDate time.Time // from go std lib
}

// we can also associate function to a struc
// by adding a receiver after the func keyword
// receiver will be in (varName *struct) which points to the struct
func (m *User) setFirstName() string {
	// this makes the function can access information from the struct
	return m.FirstName
}

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

	// numString is declared outside of this main function (available to all function in this package)
	log.Println("numString is", numString)

	// scoped to local function
	var numString2 = "eight"

	// Println() is a variadic function (can take one or more or no parameters)
	log.Println("numString2 is", numString2)

	// function that show scopes in var
	saySomethingWithSameVarNames(numString2)

	// creating a User
	user := User{
		FirstName: "Elon",
		LastName:  "Musk",
	}

	// calling the attributes of User
	// Elon Musk 0001-01-01 00:00:00 +0000 UTC
	log.Println(user.FirstName, user.LastName, user.BirthDate)

	var user2 User
	user2.FirstName = "Bill"
	user2.LastName = "Gates"

	log.Println("user is set to", user.FirstName)
	// can call the function which is part of the User struct
	// advantage is can have biz logic in the function instead of accessing the attributes directly
	log.Println("user2 is set to", user2.setFirstName())
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

func saySomethingWithSameVarNames(numString2 string) (string, string) {
	log.Println("this numString is being called in saySomethingWithSameVarName is", numString, "which is package level var")
	log.Println("this numString2 is being called in saySomethingWithSameVarName is", numString2, "which is parameter passed in function call")

	return numString2, "helllloooo"
}
