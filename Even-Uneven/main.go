package main

import (
	"fmt"
)

func main() {
	a := makeRange(0, 10)
	fmt.Println(a)

	for i := range a {
		if a[i]%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("uneven")
		}
	}

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
