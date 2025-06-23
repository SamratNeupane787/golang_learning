package main

import "fmt"
func incre(a [3]int){
	a[0]=100
}

func incrs(a []int){
	a[0]=100
}


func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3


	// b := [3]int{1, 2, 3}
	// fmt.Printf("b: %v\n",b)

	var  s []int
	s=append(s,1,2,3)
	// fmt.Printf("s: %v",s)


	incre(a)
	incrs(s)
	fmt.Printf("A: %v \n",a)

	fmt.Printf("s: %v",s)
}