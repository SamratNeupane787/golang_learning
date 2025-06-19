package main

import "fmt"


func main(){
	
	for i:=0; i<5; i++{
		fmt.Println(i)
	}

	numbers := []int{1,2,3}

	for idx, val := range numbers{
		fmt.Println(idx,val)
	}
}