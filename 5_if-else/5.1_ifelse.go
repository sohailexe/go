package main

import "fmt"

func main(){
	age:=12

	if age>= 18{
		fmt.Println("Adult")
	}else{
		fmt.Println("Not adult")
	}


	if age:=19; age>=18{
		
		fmt.Println(age)
	}else{
		fmt.Println("not adult")
	}
}