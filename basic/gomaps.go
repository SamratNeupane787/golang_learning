package main

import "fmt"

func main() {
	var m map[int]string

	if m == nil {
		fmt.Println("Map is nil")
	}

	m1 := make(map[int] string)
	m1[199]= "SAM"
	fmt.Println("%v",m1)

	val, exists := m[200]
	if !exists{
		fmt.Println("Map keyt doesnot exist")
	}else{
		fmt.Printf("map value :%v",val)
	}


	m2 := map[string]int{
		 "foo":1,
		 	"bar":2,
	}

	v1:= m2["foo"]
	fmt.Printf("%d\n",v1)
	
}