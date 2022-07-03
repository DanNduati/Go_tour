package main

import "fmt"

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

type Point struct {
	X, Y int
}

func main() {
	fmt.Println("Structs")
	fmt.Println(Vertex{1, 2})
	//struct fields are accessed using a dot.
	v := Vertex{2, 3}
	v.X = 4
	fmt.Println(v.X)
	// Pointers to structs
	// Struct fields can be accessed through a struct pointer
	p := &v
	// Access struct field
	yVal := (*p).Y
	fmt.Println(yVal)
	// Go provides a less cumbersome notation:
	p.X = 5
	fmt.Println(v)
	// Struct Literals
	// A struct Literal denotes a newly allocated struct value by listing the values of its fields
	// You can list a subset of fields by using the Name: syntax (The order of named fields is irrelevant)
	// The special prefix & returns a pointer to the struct value
	p1 := Point{1, 2} // has type Point
	p2 := Point{X: 1} // Y:0 is implicit
	p3 := Point{}     // X:0 Y:0
	ptr := &Point{2, 4}

	fmt.Println(p1, ptr, p2, p3)
}
