// package declaration by convention is main
package main

// import the fmt package from the std lib
import (
	"fmt"
	"log"
	"sort"
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

	// Map

	// the usual declaration of creating a map using make()
	// follow map[dataType of the index]dataType of the value
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	//Map can be overwritten
	myMap["dog"] = "Russell"

	// this not the conventional way
	// var myOtherMap map[string]string

	log.Println(myMap["dog"])       // print Samson to the console (overwrite by line 137 to Russell)
	log.Println(myMap["other-dog"]) // print Cassie

	// map of string int
	numMap := make(map[string]int)
	numMap["first"] = 1
	numMap["second"] = 2

	log.Println(numMap["first"])
	log.Println(numMap["second"])

	// map can hold any data type including struct
	structMap := make(map[string]User)

	randomUser := User{
		FirstName: "Jeff",
		LastName:  "Bezos",
	}

	structMap["randomUser"] = randomUser

	log.Println(structMap["randomUser"].FirstName) // shld print Jeff

	// maps are not sorted, have to look up by key
	// maps are mutable

	// slices

	// this mean a slice of strings (can put more than one thing)
	var nameSlice []string
	// adding to slice can use append(slice, newValue)
	nameSlice = append(nameSlice, "Ryan")
	nameSlice = append(nameSlice, "John")

	log.Println(nameSlice) // return [Ryan John]

	var numSlice []int

	numSlice = append(numSlice, 5)
	numSlice = append(numSlice, 0)
	numSlice = append(numSlice, 20)

	log.Println(numSlice) // return [5 0 20]

	// sorting a slice using sort.Ints() sort in increasing order by default
	sort.Ints(numSlice)

	log.Println(numSlice) // return [0 5 20]

	// shorthand to declare slices
	rangeNum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(rangeNum)

	// getting a subset of a slice using splice [firstIndex including, lastIndex excluding]
	log.Println(rangeNum[0:2]) // return [1 2]
	log.Println(rangeNum[6:9]) // return [7 8 9]

	// control flow
	// if statements
	isTrue := false // type boolean
	bigNum := 1000

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// multi conditionals
	if !isTrue && bigNum < 100 {
		log.Println("1st if-else: Both conditions satisfied")
	} else if bigNum < 2000 && isTrue {
		log.Println("2nd if-else: Both conditions satisfied")
	} else if bigNum > 1000 || isTrue {
		log.Println("3rd if-else: either condition satisfied")
	} else if bigNum == 1000 && !isTrue {
		log.Println("Last if-else: Both conditions satisfied")
	}

	// switch statements
	myAnimal := "monkey"

	// switch statements in go -> break out of the case once one condition is met (auto break out)
	switch myAnimal {
	case "cat":
		log.Println("animal is cat")
	case "dog":
		log.Println("animal is dog")
	default:
		log.Println("animal is not cat or dog")
	}

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
