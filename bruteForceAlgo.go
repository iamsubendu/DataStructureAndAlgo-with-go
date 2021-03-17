//Brute Force Algorithms are exactly what they sound like –
//straightforward methods of solving a problem that rely on
//sheer computing power and trying every possibility rather
//than advanced techniques to improve efficiency.

//For example, imagine you have a small padlock with 4 digits,
//each from 0-9. You forgot your combination, but you don't
//want to buy another padlock. Since you can't remember any
//of the digits, you have to use a brute force method to open the lock.

//So you set all the numbers back to 0 and try them one by
//one: 0001, 0002, 0003, and so on until it opens.
//In the worst case scenario, it would take 104, or
//10,000 tries to find your combination.

//to check out an element in an array

package main

import (
	"fmt"
)

//findElement method given array and k element
func findElement(arr [10]int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

// main method
func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var check bool = findElement(arr, 10)
	fmt.Println(check)
	var check2 bool = findElement(arr, 9)
	fmt.Println(check2)
}
