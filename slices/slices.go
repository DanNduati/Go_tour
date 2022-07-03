package main

import (
	"fmt"
	"strings"
)

func main() {
	// slices
	// A slice is dynammically sized, flexible view into the elements of an array
	// In practice slices are much more common than arrays
	// The type []T is a slice of type T
	// A slice is formed by specifying two indices, a low and high bound, seperated
	// by a colon
	// a[low : high]
	// This selects a half-open range which includes the first element, but excludes the last one.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
	// A slice does not store any data, it just describes a section of an underlying array
	// changing the elements of a slice modifies the corresponding elements in the corresponding array
	// Other slices that share the same underlying array will see those changes.
	names := [4]string{
		"Daniel",
		"Chege",
		"Monicah",
		"Wanjiru",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[2:4]
	fmt.Println(a, b)
	b[1] = "Nduati"
	fmt.Println(a, b)
	fmt.Println(names)
	// Slice literals
	// A slice literal is like an array literal without a length
	// This is a slice literal: []bool{true,false,true}
	q := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(q)
	r := []bool{true, false, true, true, false}
	fmt.Println(r)
	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{5, true},
	}
	fmt.Println(ss)
	// Slice defaults
	// When slicing you may omit the high and low bounds to use the defaults instead
	// The default is 0 for the low bound and the length of the slice for the high bound
	x2 := []int{2, 3, 5, 7, 11, 13}
	ss2 := x2[:2]
	ss3 := x2[4:]

	fmt.Println(ss2, ss3)
	// Slice length and capacity
	// A slice has both length and capacity
	// The length is the number of elements it contains
	// The capacity is the number of elements in the underlying array counting from the first element in the slice.
	printSlice(q)
	q = q[:0]
	printSlice(q)
	q = q[:4]
	printSlice(q)
	q = q[2:]
	printSlice(q)

	// Nil slices
	// The zero value of a slice is nil
	// A nil slice has a capacity and length of 0 and has no underlying array

	var n_slice []int
	printSlice(n_slice)
	if n_slice == nil {
		fmt.Println("Nill slice!")
	}
	// Creating a slice with make
	// This is how you create dynammically-sized arrays
	// The make function allocates a zeroed array and returns a slice that refers to that array
	t := make([]int, 5)
	printSlice(t)
	// To specify the capacity pass a third argument to make
	u := make([]int, 0, 5)
	printSlice(u)

	// slice of slices
	// slices can contain any type including other slices
	board := [][]string{
		[]string{"-", "-", "-"}, []string{"-", "-", "-"}, []string{"-", "-", "-"}, []string{"-", "-", "-"},
	}
	// The players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	//print out the board
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//Appending to a slice
	var prime_nums []int
	printSlice(prime_nums)
	//append works on nil slices
	prime_nums = append(prime_nums, 2)
	printSlice(prime_nums)
	// Add more than one element at a time
	prime_nums = append(prime_nums, 3, 5, 7, 11, 13)
	printSlice(prime_nums)
}

func printSlice(s []int) {
	fmt.Printf("Len=%d cap=%d %v\n", len(s), cap(s), s)
}
