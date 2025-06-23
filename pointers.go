package main

import "fmt"

func main(){
	var p*int
	x := 21
	p = &x
	fmt.Printf("Address of x :%d",p)
	fmt.Printf("Value of x :%d",*p)
}