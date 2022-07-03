package main

import (
	"fmt"
)

type Coord struct {
	Lat, Long float64
}

var m map[string]Coord

func main() {
	// Maps
	// A map maps keys to values
	// The zero value of a map is nil
	// A nil map has no keys, nor can keys be added
	// The make function returns a map of the given type, initialized and ready for use
	// Map literals are like struct literals but the keys are required
	m = make(map[string]Coord)
	m["Home"] = Coord{
		-0.4167, 36.9500,
	}
	location := Coord{
		-0.4167, 36.9500,
	}
	l := make(map[string]Coord)
	l["Other"] = location

	var places = map[string]Coord{
		"Home": Coord{
			-0.4167, 36.9500,
		},
		"Work": Coord{
			-1.286389, 36.817223,
		},
		"CampNou": Coord{
			41.380898, 2.122820,
		},
	}
	// if the top-level type is just a type name you can omit it from the elements of the literal

	parks := map[string]Coord{
		"National": {-1.364632, 36.832478},
		"City":     {-1.2643400390353061, 36.82816211980426},
	}

	// Mutating maps
	// Insert or update an element in a map
	parks["National"] = Coord{-1.365023, 36.849560}
	// Retrieve an element
	fmt.Println(parks["National"])
	//delete an element
	delete(places, "Work")
	// Test that a key is present with a two-value assiignment
	// If the key is not in the map then elem is the zero value of he maps element type
	elem, present := places["work"]

	fmt.Println("The value:", elem, "Present?:", present)
	fmt.Println(m)
	fmt.Println(m["Home"])
	fmt.Println(l["Other"])
	fmt.Println(places)
	fmt.Println(parks)
}
