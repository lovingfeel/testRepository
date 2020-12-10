package main

import "fmt"

func main() {
	pm := &map[string]int{"C": 1972, "Go": 2009}
	ps := &[]string{"break", "continue"}
	pa := &[4]bool{false, true, true, false}
	fmt.Printf("%T\n", pm) // *map[string]int
	fmt.Printf("%T\n", ps) // *[]string
	fmt.Printf("%T\n", pa) // *[4]bool
}