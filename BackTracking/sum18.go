// The following is an example of a backtracking algorithm.
// The problem is to identify the combinations of elements
// in an array of 10 elements whose sum is equal to 18.
// The findElementsWithSum method recursively tries to
// find the combination. Whenever the sum goes beyond the k
// target, it backtracks:

package main

import (
	"fmt"
)

//findElementsWithSum of k from arr of size
func findElementsWithSum(arr [5]int, combinations [10]int, size int, k int, addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i = l; i < size; i++ {
		//fmt.Println(" m", m)
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}

// main method
func main() {
	var arr = [5]int{10, 4, 7, 8, 9}
	var addedSum int = 18
	var combinations [10]int
	//the combination must be always greater than or eqal to addedSum
	//so that to try out new combinations
	findElementsWithSum(arr, combinations, 5, addedSum, 0, 0, 0)
	//fmt.Println(check)//var check2 bool = findElement(arr,9)
	//fmt.Println(check2)
}
