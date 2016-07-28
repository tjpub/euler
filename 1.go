package main

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

import "fmt"

const MAX int = 1000

func main() {
	fmt.Println("doing eiler 1")

	var i int = 0

	for index := 0; index < MAX; index++ {
		if index%3 == 0 || index%5 == 0 {
			i += index
		}
	}

	fmt.Printf("result: %d\n", i)
	fmt.Println("done")
}
