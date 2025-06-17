package main

import "fmt"

func main() {
	const xx = 10
	const b float64 = 3.14
	fmt.Printf("a = %v, type a = %T\nPi = %v, Type Pi = %T\n",xx,xx,b, b)
	const (
		l = 3
		w = 10
		a = l * w
	)

	fmt.Printf("Area = %v, type a = %T\n",a, a)

	const (
		one = iota
		two
		three
	)

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)

	z :=float64(a)

	fmt.Println(z)

		
}