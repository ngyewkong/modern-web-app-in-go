// test for main.go must be named main_test.go
// to check test coverage -> run "go test -cover"
// to display in an html -> run "go test -coverprofile=coverage.out && go tool cover -html=coverage.out"
package main

import "testing"

// table test in golang
// need to include the expected values as an attribute
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	// pass in the testing values
	// populate test data
	// if we want more test cases, just add it
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0, 0.0, true},
	{"expected-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

// using the table tests
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get an error")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

// manual way to do testing
// test function need to start with Test eg TestDivide
// t is the conventional name for this var which is a pointer to testing.T which is built-in package
// to run test in go, just run go test -v for verbose
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)

	if err != nil {
		// we throw a testing error here
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)

	if err == nil {
		// we throw a testing error here
		t.Error("Did not get an error when we should have as it is division by 0")
	}
}
