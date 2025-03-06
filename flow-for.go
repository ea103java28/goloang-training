package main
import "fmt"

func main(){

	var i int
for i=0; i<5; i++ {
	fmt.Println(i)
}

// 測試 1+3+5+7.....+13
var x int = 1
var result int = 0
for x <=13 {

	fmt.Print(x)
	if x < 13{
	fmt.Print("+")
	}
	
	result = result + x
	x += 2
	
}
fmt.Print("= ")
fmt.Print(result)
}