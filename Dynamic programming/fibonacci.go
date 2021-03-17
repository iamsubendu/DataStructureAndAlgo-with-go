//Fibonacci numbers computed by iteration

package main

import "fmt"

// func fibonacci(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fibonacci(n-1) + fibonacci(n-2)
// }

//we can use this
//but dynamic programmic make it more faster than
// this

func main() {
	fibonacci(7)
}

func fibonacci(n int) int {
	first := 0
	second := 1
	temp := 0
	if n == 0 {
		return first
	} else if n == 1 {
		return second
	}
	fmt.Println("0  0")
	fmt.Println("1  1")
	for i := 2; i <= n; i++ {
		temp = first + second
		first = second
		second = temp
		fmt.Println(i, "", temp)
	}
	return temp
}
