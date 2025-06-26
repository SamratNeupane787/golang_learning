package main

import "fmt"
func main() {
	numbers := []int{1, 5, 15, 35}

	for i, c := range numbers {
		fmt.Printf("%v\n %d\n", c,i)
	}
}