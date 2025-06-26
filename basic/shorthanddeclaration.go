package main

import "fmt"

func main(){
	//no need to use var keyword

	a := 3

	f :=float64(a)

	fmt.Printf("a - %d\n", a)

	arr := [3] int {1,2,3}

	fmt.Printf("arr = %v\n", arr)

	fmt.Println(f)
}