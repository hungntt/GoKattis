package main

import "fmt"

func main() {
	var a string
	var count int
	var fT, fC, fG bool
	_, _ = fmt.Scan(&a)
	for _, r := range a {
		c := string(r)
		if c == "T" && fT == false {
			fT = true
		} else if c == "C" && fC == false {
			fC = true
		} else if c == "G" && fG == false {
			fG = true
		}
		count++
	}
	if fT == fC == fG == true {
		fmt.Print(count + 7)
	} else {
		fmt.Print(count)
	}
}
