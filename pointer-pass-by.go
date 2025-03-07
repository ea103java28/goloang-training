package main
import "fmt"

func add(x int){
	x = x + 10
	fmt.Println("add function x value is ", x)
}

func add2(x *int){
	*x = *x + 10
	fmt.Println("add2 function x value is ", *x)
}


func main(){
	var x int = 10
	add(x)
	fmt.Println("pass by value x value is ", x)

	add2(&x)
	fmt.Println("pass by pointer x value is ", x)


}