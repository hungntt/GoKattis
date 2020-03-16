package main

import "fmt"

func main() {
	var l, r int
	_, _ = fmt.Scan(&l, &r)
	if l == r {
		if l == 0 {
			fmt.Print("Not a moose")
		} else {
			fmt.Print("Even ", l*2)
		}
	} else if l > r {
		fmt.Print("Odd ", l*2)
	} else {
		fmt.Print("Odd ", r*2)
	}
}
