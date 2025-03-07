package main
import  "fmt"

func main(){
	slice := []int{1, 2, 3} // slice 不需要固定長度，動態長度
	slice = append(slice, 4, 5) // slice can append
	fmt.Println("slice can append") 
	for _, value := range slice {
		fmt.Println(value) 
	}

	arr := [3]int{1, 2, 3} // arr 會固定長度
	fmt.Println("arr can't append") 
	for _, value := range arr {
		fmt.Println(value) 
	}
}