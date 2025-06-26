package main

import "fmt"

func main(){
	day :=3
	 
	switch day{
	case 1:
		fmt.Println("Sunday")
	
	case 2:
		fmt.Println("Monday")
	
	case 3:
		fmt.Println("Tuesday")

	default: 
	fmt.Println("other day")


}

//multiple cases values 

switch day{
case 1,2,3:
	fmt.Println("start of the week")


case 4,5:
	fmt.Println("end of the week")


default:
	fmt.Println("weekedn")}
}