//The Tower of Hanoi (also called the Tower of Brahma) We are given three
//rods and N number of disks, initially all the disks are added to first rod (the
//leftmost one) in decreasing size order. The objective is to transfer the entire
//stack of disks from first tower to third tower (the rightmost one), moving only
//one disk at a time and never a larger one onto a smaller.

//in short we have to shift the discs from left to right..
//only one disc can be moved at a time and no larger dic will come
//on top of smaller

//use simple recursion
package main

import (
	"fmt"
)

func main() {
	TowersOfHanoi(3)
	//3 disks
}
func DiscManagement(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	DiscManagement(num-1, from, temp, to) //recursive function
	fmt.Println("Move disk ", num, "kg from peg ", from, " to peg ", to)
	DiscManagement(num-1, temp, to, from)
}
func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are:")
	DiscManagement(num, "A", "C", "B")
}

//recursive function is used in this so it is also called recursive algorithm
