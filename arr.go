package main
import "fmt"

func main(){
	var arr1 [3]int 
	var arr2 [4]string 
	var arr3 [5]bool

	fmt.Println("沒給初始值，會給預設資料")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	 arr1 = [3]int {100, 99 ,88}
	 arr2 = [4]string {"x", "y", "z"}
	 arr3 = [5]bool {true, true, true}

	fmt.Println("給初始值")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	var arr4 [3] int
	arr4[0] = -99
	arr4[1] = -99
	arr4[2] = -99
	fmt.Println("逐一給值")
	fmt.Println(arr4)

	showArr(arr4[:])

}

func showArr (arr []int ){
	var i int
	for i=0; i<len(arr); i++ {
		fmt.Println(arr[i])
	}
}