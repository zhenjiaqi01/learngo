package main

import "fmt"

func main() {
	p1 := testPtr1()
	fmt.Println("p1 ", p1)
	fmt.Println("*p1 ", *p1)

}

//定义一个函数，返回一个string类型的指针
func testPtr1() *string {
	city := "深圳"
	ptr := &city
	return ptr
}
