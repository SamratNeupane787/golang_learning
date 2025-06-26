package main

import "fmt"

type Car struct {
	Name  string
	Model string
	Year  int
	Brand string
	Operation int
}

//method : years of operation, whetherit is operational or not

func (c Car) Operational(){
		if c.Operation >15{
			fmt.Print("Car is not operational")
		}else{
			fmt.Print("Car is operational")
		}
}

func main() {
	c := Car{"Tesla Model S", "Model S", 2025, "Tesla",2}
	fmt.Print(c)
	c.Operational()
}