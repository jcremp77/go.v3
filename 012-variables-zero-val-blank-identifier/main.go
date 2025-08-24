package main

import "fmt"

func main() {
	// short declaration operator `:=` to assign value of int 42 to 'a'
	a := 42
	fmt.Println(a)
	// the int value of '3' is assigned to a blank identifier
	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	// the following code works because '_' does not need to be called
	fmt.Println(b, c, d, f)

	// this would not work becuase 'e' is not declared in the fmt.Println() statement
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	// declare variable 'g' of type INTEGER with no value
	var g int
	// following statement will print out '0'
	fmt.Println(g)
	// integer value of '43' is assigned to variable 'g'
	g = 43
	// the integer value of '43' is printed out
	fmt.Println(g)

	// declare a variable to hold a VALUE of a certain TYPE
	// initializing a variable
	var h int = 44
	fmt.Println(h)

}
