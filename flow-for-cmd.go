package main
import "fmt"

func main(){

	fmt.Println("break")

	var x int = 0
	for x < 5 {
		if x >= 3 {
			break
		}
		fmt.Println(x)
		x++
	}

	fmt.Println("continue")
	for x=0; x<=5; x++ {
		
		if x %2 == 0 {
			continue
		}
		fmt.Println(x)
		
	}

	var result int = 0
	for true {
		fmt.Println("請輸入數字，直到你輸入0 會結束")
		var n int 
		fmt.Scanln(&n)
		if n==0 {
			break
		}
		result += n
	}
	fmt.Println("result is: " , result)

}