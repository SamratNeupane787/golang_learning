package main

import "fmt"

func main() {
	
	type Power struct{
		HP string

	}

	type Car struct {
		Name  string
		Type  string
		Brand string
		years int
		Power 
	}


	// c := Car{
	// 	Name:  "Model s",
	// 	Type:  "Ev",
	// 	Brand: "Tesla",
	// 	years: 4,
	// 	Power{
	// 			HP:"319",
	// }
		
	// }

	c := Car{ "MODEL Y","EV","TESLA",3,Power{"310"}}
	fmt.Println(c)
	
}