package main

import (
	"fmt"
)

//basic structure

// func functionName (p1, p2)(r1,r2){
// 	//funtion body
// 	//return statement
// }

//variadic funtions -- used when you dont now how many parameter you are going to pass to the function
//  func sum(numbers ...int) int {
// 	//total sum
// 	//return total
//  }
// example:- sum(1,2,3) sum(1,2,3,4,5,6)

//anonymous functions

// great := func(name string){
// 	 	fmt.PrintLn("HELLO", name)
// }

// greet("SAMRAT")

// function closures

// -special types of anonymous function
// -references variables outside of its sccope

func Hello( a int , b int) int{
	fmt.Printf("Hello people! the sum of two vairbles is: %v\n", a+b)
	return a+b
}

// func main(){
// 	ans := Hello(1,2)
// 	fmt.Println(ans)
// }

func Add(number ...int)int{
	total :=0

	for i:= range number{
		total += number[i]
	}
	return total
}

func main(){
	greet :=func (name string)  {
		fmt.Println("Hello , from GFG ", name)
		
	}

	greet("sam")
}