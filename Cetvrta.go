package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	p1 := Point{}
	p2 := Point{}
	p3 := Point{}
	p4 := Point{}
	_, _ = fmt.Scan(&p1.x, &p1.y)
	_, _ = fmt.Scan(&p2.x, &p2.y)
	_, _ = fmt.Scan(&p3.x, &p3.y)

	if p1.x == p2.x {
		if p3.y == p2.y {
			p4.x = p3.x
			p4.y = p1.y
		} else {
			p4.x = p3.x
			p4.y = p2.y
		}
	} else if p2.x == p3.x {
		if p1.y == p2.y {
			p4.x = p1.x
			p4.y = p3.y
		} else {
			p4.x = p1.x
			p4.y = p2.y
		}
	} else if p1.x == p3.x {
		if p2.y == p3.y {
			p4.x = p2.x
			p4.y = p1.x
		}
	} else {
		p4.x = p2.x
		p4.y = p3.y
	}
	fmt.Print(p4.x, p4.y)
}
