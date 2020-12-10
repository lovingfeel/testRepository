package main

import (
	"fmt"
)

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

//func main()  {
//	var a = []int{1,2,3,4,5}
//	var slice = make([]int ,12,13)
//	fmt.Println(slice)
//	copy(slice,a)
//	fmt.Println(slice)
//}
