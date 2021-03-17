//A  queue is a First-In-First-Out (FIFO) kind of data structure.
//The element that is added to the queue first will be the
//first to be removed and so on.

// Queue ADT Operations:
// · Add(): Add a new element to the back of the queue.
// · Remove(): remove an element from the front of the queue
// and return itsvalue.
// · Front(): return the value of the element at the front
// of the queue.
// · Size(): returns the number of elements inside the queue.
// · IsEmpty(): returns 1 if the queue is empty otherwise
// return 0.

package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func main() {
	q := queue.New()
	q.Enqueue(1)
	fmt.Println(q)
	q.Enqueue(2)
	fmt.Println(q)
	q.Enqueue(3)
	fmt.Println(q)
	q.Enqueue(4)
	fmt.Println(q)

	for q.Len() != 0 {
		val := q.Dequeue()
		fmt.Print(val, " ")
	}
}
