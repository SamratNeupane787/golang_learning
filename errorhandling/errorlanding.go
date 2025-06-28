package main

import (
	"errors"
	"fmt"
)

func check(num int) (string, error) {
	if num < 0 {
		return "", errors.New("Number is negetive")
	}
	return  "Number is postitive",nil
}

func read(filename string ) error{
	if filename == ""{
		return errors.New("file is empty")
	}
	return  nil
}

func file()(string ,error){
	e := read("")
	if e!= nil{
		return  "", fmt.Errorf("Error reading file:%w",e)
	}
	return "file exist",nil
}

func main(){

	_,e :=check(-1)
	if e!=nil{
		fmt.Println(e)
	}

	_,err := file()

	if err != nil{
		fmt.Println(err)
		we := errors.Unwrap(err)

		fmt.Println(we)
	}

}