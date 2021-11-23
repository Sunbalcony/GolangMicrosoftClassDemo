package main

import "fmt"

type triangle struct {
	size  int
	size2 int
}
type coloredTriangle struct {
	triangle
	color string
}

func (t *triangle) perimeter() int {
	t.size2 = t.size * 3
	return t.size2

}
func (t *triangle) print()  {
	fmt.Println(t)

}
func main() {
	t := triangle{size: 2}
	fmt.Println(t)
	fmt.Println("perimeter", t.perimeter())
	//fmt.Println(t)
	t.print()
	//t2:=coloredTriangle{triangle{size: 6},"blue"}
	//fmt.Println("size is",t2.size)
	//fmt.Println("Perimeter",t2.perimeter())

}
