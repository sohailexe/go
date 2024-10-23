package main


func sum(nums ...int) int {
	total:=0;
	for _,num:= range nums {
		total+=num;
	}
	return total;
}
func main(){
println(sum(2,2,3))
}