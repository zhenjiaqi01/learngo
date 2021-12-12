package main

import "fmt"

//go语言中没有枚举类型，可以使用cont + iota（常量累加器）来进行模拟
//模拟一周的枚举
//iota 常量累加器 是从零开始 每换一行递增一
//常量组有个特点 如果不赋值默认与上一行表达式相同
//如果同一行出现两个iota 这两个值是相同的  const 常量属于预编译，所以不需要：= 自动推导
//每个常量组的iota是独立的 如果遇到const iota 会清零
const (
	MONDAY    = iota // 0
	TUESDAY   = iota // 1
	WEDNESDAY = iota
	THURSDAY  = iota
	FRIDAY    = iota
	SATURDAY  = iota
	SUNDAY
)
const (
	JANU = iota + 1
	FEB
)

func main() {
	//var number int
	//var name string
	//var flag bool
	//
	////变量组统一定义变量
	//var (
	//	number int
	//	name string
	//	flag bool
	//)
	fmt.Println("MONDAY is ", MONDAY)
	fmt.Println(JANU)

}
