package main

import "fmt"

func main() {
	//go语言使用指针会使用内部的垃圾回收机制gc，开发人员不需要手动释放内存
	name := "lily"
	ptr := &name

	fmt.Println("name ", *ptr)
	fmt.Println("name address", ptr)

	//new 关键字定义指针
	name2Ptr := new(string)
	*name2Ptr = "Duke"

	//可以返回栈上指针，编译器会自动判断这段代码，将city分配到堆内存上
	fmt.Println("name2", *name2Ptr)
	fmt.Println("name2 ptr", name2Ptr)

	res := testPtr()
	fmt.Println("res ", res)

}

//定义一个函数，返回一个string类型的指针
func testPtr() *string {
	city := "深圳"
	ptr := &city

	return ptr
}
