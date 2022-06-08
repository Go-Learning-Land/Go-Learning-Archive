package main

import "testing"

/*
It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
*/
func TestHello(t *testing.T) { //The test function takes one argument only t *testing.T
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want) //Log + Fail ==> Errorf
		/*
			We are calling the Errorf method on our t which will print out a message and fail the test.
			The f stands for format which allows us to build a string with values inserted into the placeholder values %q.
			When you made the test fail it should be clear how it works.

		*/
	}
}

//Test for Hello function
