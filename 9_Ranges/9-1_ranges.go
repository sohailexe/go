package main

import "fmt"

// iterating over datastructures
func main(){
	nums:= []int{1,2,3,4,5}
	fmt.Println(nums)

	for index,num := range nums{
		fmt.Println(num,"         index" ,index)
	}

	m:= map[string]string{"name":"Sohail", "age":"22 years"}

	for key,val := range m{
		fmt.Println(key, "          ",val)
	}
}