package main

import "fmt"

func change(m map[string]int){
	 m["apple"]=50
}

func main(){
	m:=map[string]int{
		"apple":1,
		"orange":2,
	}
	fmt.Print(m)

	fmt.Print(m["apple"])

	if val,ok :=m["apple"];ok{
		fmt.Println("old value",val)
		m["apple"]=15
	}
	fmt.Println(m["apple"])

	delete(m,"apple")
	change(m)
	fmt.Println(m["apple"])


}