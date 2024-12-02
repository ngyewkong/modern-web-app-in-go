package helpers

import "math/rand"

type SomeType struct {
	TypeName   string
	TypeNumber int
}

// create a random number function
func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}
