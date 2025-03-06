package main
import "fmt"


func test (msg string){
	fmt.Println(msg)
}

func sum (max int){
	var i int 
	var result int = 0
	for i=1; i<=max; i++ {
		result += i
	}
	fmt.Println(result)
}

func main(){
	test("xxxxxxxxxxxx")
	sum(10) //1+2+3...+10
	sum(50) //1+2+3+4...+50
	sum(100)  //1+2+3+4+....+100
}



