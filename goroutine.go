package main
import (
	"fmt" 
	"time"
)

func main(){
	fmt.Println("開始找質數")
	num := 3000000
	start := time.Now()
	for i := 1; i<=num; i++{
		if isPrime(i){
			fmt.Println(i)
		}
	}
	end := time.Now()
	fmt.Println(end.Unix() - start.Unix(), "seconds")
}

func isPrime(num int) bool { // 找質數
	if num == 1{
		return false
	}else if num == 2 {
		return true
	}else {
		for i :=2; i<= num; i++ {
			if num %i == 0{
				return false
			}
		}
		return true
	}
}

