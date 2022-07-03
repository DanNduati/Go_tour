package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// Range
	// The range form of the for loop iterates over a slice or map
	// When ranging over a slice, two values are returned for each iteration
	// Index and the copy of the element at that Index
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
	// You can skip the index or value by assigning it to _
	// for _,value := range pow
	// for i,_ := range pow
	for _, elem := range pow {
		fmt.Println(elem)
	}
	// if you only want the index you can completely omit the second variable
	// for i:=range pow
	sqr := make([]int, 10)
	for i := range sqr {
		sqr[i] = 1 << uint(i) // 2**i
	}
	for _, value := range sqr {
		fmt.Printf("%d\n", value)
	}
}
