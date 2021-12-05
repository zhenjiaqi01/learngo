package main

import "fmt"

//函数返回值在参数列表之后的，如果有多个返回值需要用圆括号包裹
func test(a int, b int, c string) (int, string, bool) {

	return a + b, c, true //多个返回值
}

//如果只有一个返回值，并且没有名字 那么不需要加圆括号
func test3(a, b int, c string) (res int, str string, b1 bool) {
	//直接使用返回值的变量名参与运算
	res = a + b
	str = c
	b1 = true

	//当使用了有名字的返回值时，可以直接直接写return
	return

}
func main() {
	v1, s1, _ := test(10, 20, "Hello")
	fmt.Println("v1 ", v1)
	fmt.Println("s1 ", s1)

}
