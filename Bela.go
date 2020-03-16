package main

import "fmt"

func main() {
	var hand, sum int
	var suit string
	_, _ = fmt.Scan(&hand, &suit)
	for i := 0; i < hand*4; i++ {
		var str, num, numsuit string
		_, _ = fmt.Scan(&str)
		num = string(str[0])
		numsuit = string(str[1])
		switch num {
		case "A":
			sum += 11
		case "K":
			sum += 4
		case "Q":
			sum += 3
		case "J":
			if numsuit == suit {
				sum += 20
			} else {
				sum += 2
			}
		case "T":
			sum += 10
		case "9":
			if numsuit == suit {
				sum += 14
			}
		}
	}
	fmt.Print(sum)
}
