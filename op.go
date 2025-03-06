package main
import "fmt"
func main(){

// 1-1. 算數運算 +, -, *, /
var x int
x = 3 + 3 
fmt.Println(x)

// 1-2. 指定運算 =, +=, -=, *=, /=
 x += 4 // x = x + 4
fmt.Println(x)

// 1-3. 單元運算 ++, --
x ++
fmt.Println(x)
x --
fmt.Println(x)

// 1-4. 比較運算 >, <, >=, <=, ==

var result bool = 4 > 3
fmt.Println(result)

// 1-5. 邏輯運算 !, &&, ||

result = 4 > 3 || 4 < 3
fmt.Println(result)

result = 4 > 3 && 4 < 3
fmt.Println(result)

}

