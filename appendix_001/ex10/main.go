package main

import (
	"fmt"
)

func main() {

	xs := []int{28, 25, 29, 11, 12, 13, 22, 17, 1}

	h := 3

	for i := 1; i < len(xs); i++ {
		for j := 0; j <= i; j++ {
			if xs[i] < xs[j] {
				for k := i; k >= j+1; k-- {
					xs[k], xs[k-1] = xs[k-1], xs[k]
				}
			}
		}
	}

	fmt.Println(xs)

}
