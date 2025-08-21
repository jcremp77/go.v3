// This program provides an example using Print Formatting

package main

import (
	"fmt"
)

func who() {
	const name, age = "Ethan", 23
	fmt.Printf("%s is %d years old.\n", name, age)
}

var _ = who
