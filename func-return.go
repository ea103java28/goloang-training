package main
import "fmt"

func show (msg string){

	if(msg == "xxx"){
		return // 即便沒有回傳值，也會強制結束方法
	}
	fmt.Println(msg)
}

func add (n1 int, n2 int) int {
	var result int = n1 + n2
	return result
}

func main(){
	show("test")
	show("test2")
	show("xxx")
	show("test3")

	fmt.Println(add(5, 5))
	fmt.Println(add(10, 10))
}
