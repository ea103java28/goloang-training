package main
import "fmt"

func main(){
	var money int
	fmt.Println("請輸入要領多少$$")
	fmt.Scanln(&money)
	
	if money <= 100000{
		fmt.Println("OK")
	}else {
		fmt.Println("too much")
	}
}