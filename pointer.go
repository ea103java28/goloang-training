package main
import "fmt"

/*
建立指標變數 > 取得記憶體位址
存放到指標變數 > 反解指標變數

原始指標型態: int、float64、string
原始指標型態: *int、*float64、*string
*/


func main(){

// 取得記憶體位址: &變數名稱
var x int = 3
fmt.Println(&x)

// 存放到指標變數
// 指標資料型態: *資料型態
var xPtr *int = &x
fmt.Println(xPtr)

//反解指標變數
// 反解運算: *指標變數名稱
fmt.Println(*xPtr)
fmt.Println(*&x)


var word string = "hello"
var wordPtr *string = &word
fmt.Println(&word)
fmt.Println(wordPtr)
//反解指標
fmt.Println(*&word)
fmt.Println(*wordPtr)


}