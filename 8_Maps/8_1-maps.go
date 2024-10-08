package main

import "fmt"

func main(){
	// creating map
	

	m:= make(map[string]string)
	m["name"]= "sohail"
	m["age"]= "22-years old"
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "age")
	fmt.Println(len(m))

	
}