package main
import (
	"fmt"
)

func subtactOne(numbers []int)  []int {
	for i := range numbers {
		numbers[i] -= 2
	}
	numbers=append(numbers,12)
	return numbers
}

func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	bi:=subtactOne(nos)                               // function modifies the slice
	fmt.Println("slice after function call", bi) // modifications are visible outside
}
