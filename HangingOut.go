package main

import (
	"fmt"
)

func main() {
	var max, line, current, warning int
	_, _ = fmt.Scan(&max, &line)
	for i := 0; i < line; i++ {
		var status string
		var num, temp int
		_, _ = fmt.Scan(&status, &num)
		if status == "enter" {
			temp = current + num
			if temp <= max {
				current += num
			} else {
				warning++
			}
		} else {
			current -= num
		}
	}

	fmt.Println(warning)
}
