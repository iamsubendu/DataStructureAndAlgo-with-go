package main

import (
	"fmt"
)

func make3D(m, n, p int) [][][]int {
	buf := make([]int, m*n*p)

	x := make([][][]int, m)
	for i := range x {
		x[i] = make([][]int, n)
		for j := range x[i] {
			x[i][j] = buf[:p:p]
			buf = buf[p:]
		}
	}
	return x
}

func main() {
	fmt.Println(make3D(2, 4, 5))
}
