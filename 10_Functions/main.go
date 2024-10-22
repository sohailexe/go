package main

func add(a int, b int) int {
	return a+b
}

func languages() (string,string,string){
	return "Urdu", "english", "Norvigen";
}
func main(){
	lang1,lang2,lang3:= languages()

	
	println(lang1,lang2,lang3)
	println(add(2,4))
}