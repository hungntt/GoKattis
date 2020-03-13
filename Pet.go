package main

import (
	"fmt"
)

func main() {
	best := 0
	index := -1
	for i := 0; i < 5; i++ {
		var a, b, c, d int
		_, _ = fmt.Scan(&a, &b, &c, &d)

		if a+b+c+d > best {
			best = a + b + c + d
			index = i + 1
		}
	}
	fmt.Print(index, best)
}
