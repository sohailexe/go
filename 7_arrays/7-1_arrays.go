package main

import "fmt"

func main(){
	var nums[4] int
	nums[0]=1
	nums[1]=2

	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)

	newNums := [4]int{1,2,3,4}
	fmt.Println(newNums)

	twoD := [2][2]int {{1,2},{3,4}}
	fmt.Println(twoD)
}