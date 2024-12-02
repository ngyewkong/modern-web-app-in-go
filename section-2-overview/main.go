// package declaration by convention is main
package main

// creating out own packages
// good practice to use the github name of the proj (in case you push to github later)
// go mod init github.com/username/nameOfPackage
// to use a custom package just need to import

// import the fmt package from the std lib
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/ngyewkong/myniceprogram/helpers"
	"golang.org/x/exp/rand"
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

// interfaces in golang
type Animal interface {
	// inside this interface definition is the list of functions that any var that satisfies the animal interface must have
	Says() string
	NumberOfLegs() int
}

// two custom types DOg & Gorilla
type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

// create a struct to handle the attributes
type Person struct {
	// special tag to handle/receive json (json:key_name_of_json)
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
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

	// looping and ranging over data
	// only for loop available in golang
	// print out to console 0 to 10
	for i := 0; i <= 10; i++ {
		log.Println("i is:", i)
	}

	animals := []string{"chat", "chien", "cheval", "poisson", "vanche"}

	// iterate over the string slice using range
	// give two arguments (i & animal) -> i is the current iteration & the second argument is the value
	for i, animal := range animals {
		log.Println(i, animal) // print out 0 chat 1 chien ...
	}

	// but if we only need the value but not the iteration
	// we cannot NOT USE i in the loop as golang do not compile as we have unused variables
	// we have to use the underscore _ (blank identifier)
	for _, animal := range animals {
		log.Println(animal) // print out chat chien cheval poisson vanche ...
	}

	// range function can range over slices, maps, even strings
	pets := make(map[string]string)
	pets["cat"] = "Garfield"
	pets["dog"] = "Carrefour"

	// this will iterate the pets map, which is a map of string (key) string (value)
	// this will print out the keys & values of the map
	for petType, petName := range pets {
		log.Println(petType, petName)
	}

	// range over string
	var firstLine = "Once upon a time, there is ..."

	// the values printed is a byte (golang doc)
	// string in golang is actually a slice of bytes or runes
	for i, char := range firstLine {
		log.Println(i, ":", char)
	}

	// strings are actually immutable
	// so reassignment is actually complex (involving destroying existing object in memory and creating new object)
	firstLine = "x"
	for i, char := range firstLine {
		log.Println(i, ":", char)
	}

	// we can iterate over custom objects
	type EmailUser struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var emailUsers []EmailUser
	emailUsers = append(emailUsers, EmailUser{"Jeff", "Bezos", "jeffbezos@amzn.com", 45})
	emailUsers = append(emailUsers, EmailUser{"Tim", "Cook", "timcook@apple.com", 55})
	emailUsers = append(emailUsers, EmailUser{"Jensen", "Huang", "jensenhuang@nvda.com", 50})
	emailUsers = append(emailUsers, EmailUser{"Donald", "Trump", "donaldjtrump@trump.com", 65})

	// print out the values of each element in emailUsers array that is a custom struct EmailUser
	for _, l := range emailUsers {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

	// demo interface
	dog := Dog{
		Name:  "Samson",
		Breed: "Golden Retriever",
	}

	// invoke PrintInfo() function that is using the Animal Interface
	// setting PrintInfo(dog) -> compiler error
	// cannot use dog (variable of type Dog) as Animal value in argument to PrintInfo: Dog does not implement Animal (missing method NumberOfLegs)
	// need to create functions for Dog Struct that implements the two required methods for Animal Interface
	// &dog passes the reference to dog instead as the receiver are pointer types
	PrintInfo(&dog) // print "This animal says Woof and has 4 legs"

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 38,
	}

	// pass by reference
	PrintInfo(&gorilla) // This animal says OOOFFFFF and has 2 legs

	// TYPES that need pointers -> (slices, maps, functions)
	// TYPES that do not need pointers -> (Strings, Ints, Floats, Booleans, Arrays, Structs)

	// using custom package
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)

	// channels in golang
	// create channels using make() function
	// chan int -> creating an integer channel
	// channel is a place to send information which will be received in one or more places in the program
	intChan := make(chan int)

	// good practice to defer and close the int chan
	// defer close(intChan) -> after the keyword defer, execute that as soon as the current function is done
	// useful in opening files, connecting to db
	// leaving all files open during read/write -> will run out of handles at a given time
	defer close(intChan)

	// running the routine as its goroutine (concurrent) using go keyword
	go CalculateValue(intChan)

	// getting the CalculateValue randomNumber from the int chan
	randNum := <-intChan
	log.Println(randNum) // printing random numbers 0 to 999

	// reading & writing JSON
	// note the JSON cannot have unnecessary commas (end of last k-v pair, end of last entry)
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]
	`
	// unmarshall JSON using the custom Person struct that we created
	// is a slice of Person as this JSON might have one or more entries
	var unmarshalled []Person

	// using json.Unmarshal()
	// takes a slice of bytes, an interface that you going to put the contents into
	// to convert a string to a slice of bytes []byte(myJson)
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	// %v refers to an interface, which is unmarshalled
	log.Printf("unmarshalled: %v", unmarshalled) // unmarshalled: [{Clark Kent black true} {Bruce Wayne black false}]

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Logan"
	m1.LastName = "Wolverine"
	m1.HairColor = "black"
	m1.HasDog = false

	var m2 Person
	m2.FirstName = "Scarlett"
	m2.LastName = "Witch"
	m2.HairColor = "red"
	m2.HasDog = true

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	// in production use json.Marshal (very small json but not readable)
	// json.mMarshalIndent (make it readable with indentation)
	// using no prefix
	newJson, err := json.MarshalIndent(mySlice, "", "	")
	if err != nil {
		log.Println("error marshalling", err)
	}
	// normal write will be a slice of bytes (that is what is being returned from json.Marshall)
	// use fmt package to format and convert the slice of bytes to string
	fmt.Println(string(newJson))
	// [
	//
	//	        {
	//				"first_name": "Logan",
	//				"last_name": "Wolverine",
	//				"hair_color": "black",
	//				"has_dog": false
	//		},
	//		{
	//				"first_name": "Scarlett",
	//				"last_name": "Witch",
	//				"hair_color": "red",
	//				"has_dog": true
	//		}
	//
	// ]

	// call divide function
	result, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of division is", result)
}

const numPool = 1000

// function that takes in an channel of integers
func CalculateValue(intChan chan int) {
	// as of go >1.20, no need to specify seed number unless you want reproducibility
	rand.Seed(uint64(time.Now().UnixNano()))
	randomNum := helpers.RandomNumber(numPool)
	// pass the randomNum into the int chan
	intChan <- randomNum
}

// utility function that use the Animal Interface
func PrintInfo(a Animal) {
	// calls the available methods from the interface
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

// to make the Dog struct type satisfy the Animal interface
// need to implement the two required functions of the interfaces for Dog struct
// func (receiver structType) requiredMethod() returnType
// in golang usually receivers are pointer types as it just make things faster (*Dog)
func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "OOOFFFFF"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
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

// for demo-ing testing in go
// error is a built in interface
func divide(x, y float32) (float32, error) {
	var result float32

	// raise error when division by 0
	if y == 0 {
		// errors.New() allow us to create our own error message
		return result, errors.New("cannot divide by 0")
	}

	result = x / y
	// we can return nil to error as the base case is handled
	return result, nil
}
