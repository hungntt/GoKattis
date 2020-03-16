package main

import "fmt"

func main() {
	ret := 1
	var a string
	_, _ = fmt.Scan(&a)
	for _, r := range a {
		c := string(r)
		if c == "A" {
			if ret == 1 {
				ret++
			} else if ret == 2 {
				ret--
			}
		} else if c == "B" {
			if ret == 2 {
				ret++
			} else if ret == 3 {
				ret--
			}
		} else {
			if ret == 3 {
				ret -= 2
			} else if ret == 1 {
				ret += 2
			}
		}
	}
	fmt.Print(ret)
}
