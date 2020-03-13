package main

import "fmt"

func main() {
	var n, h, v int
	_, _ = fmt.Scan(&n, &h, &v)
	arr := []int{
		h * v * 4,
		h * (n - v) * 4,
		(n - h) * (n - v) * 4,
		(n - h) * v * 4,
	}
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value // found another smaller value, replace previous value in max
		}
	}

	fmt.Println(max)

}
