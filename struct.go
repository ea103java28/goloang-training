package main
import "fmt"

type Point struct{
	x int
	y int
}

func main(){
	var p1 Point = Point{1, 2}
	var p2 Point = Point{y:1, x:2}
	fmt.Println(p1.x, p1.y)
	fmt.Println(p2.x, p2.y)
	p2.x = 100
	fmt.Println(p2.x, p2.y)
}