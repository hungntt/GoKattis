package main

import "fmt"

func main() {
	var line int
	for {
		fmt.Scanln(&line)
		if line == -1 {
			break
		}
		var prev_time, res int

		for i := 0; i < line; i++ {
			var speed, time, duration int
			_, _ = fmt.Scanln(&speed, &time)
			duration = time - prev_time
			res += speed * duration
			prev_time = time
		}
		fmt.Println(res, "miles")
	}
}