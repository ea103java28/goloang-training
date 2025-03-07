package main
import (
	"fmt" 
	"time"
)

func main(){
	fmt.Println("開始找質數.....")
	num := 90000
	counter := 0
	start := time.Now()
	slice := []int{}

	for i := 1; i<=num; i++ {
		if isPrime(i){
			// fmt.Println(i)
			slice = append(slice, i)
			counter++
		}
	}
	end := time.Now()
	fmt.Println("no use goroutine need", end.Unix() - start.Unix(), "seconds")
	fmt.Println("total primes is ", counter)

}

func isPrime(num int) bool { // 找質數
	if num == 1 {
		return false
	}else if num == 2 {
		return true
	}else {
		for i := 2; i < num; i++ {
			if num % i == 0{
				return false
			}
		}
		return true
	}
}

