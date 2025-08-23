// This program is part of package named 'utils'
// and provides an example using Print Formatting

package utils

import (
	"fmt"
)

func NameAge() {
	const name, age = "Ethan", 23
	fmt.Printf("%s is %d years old.\n", name, age)
}

var _ = NameAge
