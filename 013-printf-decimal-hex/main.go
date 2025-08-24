package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	// short declaration only within a func() body, E.g., between {}'s
	adams := 42

	fmt.Println()

	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	fmt.Println()
	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "N\tTYPE\tBINARY\tHEX")
	for _, v := range []int{a, b, c, d, e, f} {
		fmt.Fprintf(w, "%d\t%T\t%b\t%x\t\n", v, v, v, v)
	}
	w.Flush()
}
