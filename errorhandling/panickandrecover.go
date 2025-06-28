package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("Divison by zero")
	}

	return a / b

}


func main() {
	

	defer func(){
		if r:= recover(); r!= nil{
			fmt.Println("Recovered form panic : %v" ,r)
		}
	}()

	z := divide(4, 0)
	fmt.Println(z)

}