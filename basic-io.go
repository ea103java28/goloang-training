package main
import "fmt"
func main(){
	var x int
	var y int
	fmt.Println("請輸入兩個數字，會回傳相加結果")
	fmt.Scanln(&x, &y)// &變數名稱:取得變數的指標(pointer)
	fmt.Println(x + y)
}