package main

import "fmt"

type CustomError struct{
	Code int
	Message string
}

func (e CustomError)Error() string{
	return  fmt.Sprint("Error code %d with message : %s", e.Code, e.Message)
}

func check(code int)(string,error){
	if code == 400{
		return "", &CustomError{Code:400, Message: "Invalid request"}
	}
	return "'success'", nil

}
func main(){
	if _, error :=check(400); error!=nil{
		fmt.Println(error)
	}


}