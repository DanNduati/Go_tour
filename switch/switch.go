package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("The rest!")
	}
	fmt.Print("When's Friday? : ")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
	// Switch with no condition
	// This construct is a clean way to write long if-then-else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
	// Defer -> defers the execution of a function until the surrounding function returns
	defer fmt.Println("World!")
	fmt.Print("Hello ")
	// Stacking defers
	// Deferred functions are pushed ontu the stack.
	// The deffered calls are executed in the last in first out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println("The value is: ", i)
	}
	fmt.Println("Done")
}
