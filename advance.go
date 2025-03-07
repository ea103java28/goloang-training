package main
import  f "fmt"

func main(){
	/*
	:=（短變數宣告）
	:= 是 Go 獨有的語法，用於短變數宣告（short variable declaration）。
	它的作用是 自動推斷變數型別 並宣告並初始化變數。
	只能用在函式內部，不能用來宣告全域變數。	
	*/
	 x := 10
	 str := "hello"
	 f.Println(x, str)


	arr := []int{1, 2, 3, 4, 5}

	// 使用 `_` 忽略索引
	for _, value := range arr {
		f.Println(value) // 只會輸出值
	}

	/*
	range arr 會遍歷 arr（這是一個 slice 或陣列）。
	這裡 range arr 會產生兩個值：
	index（索引）：arr 中當前元素的索引（從 0 開始）。
	value（值）：arr 中當前索引位置的值。
	*/
	for index, value := range arr {
		f.Println(index, value) // index, value
	}


}