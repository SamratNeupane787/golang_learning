package main

import "fmt"

func main() {
	s := "hello!😂"

	r := []rune(s)

	fmt.Println(len(s))
	fmt.Println(len(r))


	 for i, c:=range s{
		fmt.Printf("String character: %c and index : %d with ascii value : %v\n",c,i,c)
	 }
	 for i, c:=range r{
		fmt.Printf("String character: %c and index : %d with ascii value : %u\n",c,i,c)
	 }
}
