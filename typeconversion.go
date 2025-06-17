package main

import "fmt"

func main() {
	var i int = 3
	f := float64(i)

	fmt.Printf("My float var :%f",f)

	var fl float64 = 3.14

	in := int(fl)

	fmt.Printf("My in var :%d",in)

	var bt []byte = []byte{99,100,111}
	st := string(bt)
	fmt.Printf("My string : %v\n", st)
	
}