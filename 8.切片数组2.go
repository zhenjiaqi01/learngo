package main

import "fmt"

func main() {
	names := [7]string{"北京", "上海", "广州", "深圳", "洛阳", "南京", "秦皇岛"}

	//基于names创建一个新的数组
	//names1 :=[]string{}
	names1 := names[0:3] //左闭右开区间
	names1[1] = "天津"     //修改names1 names的值也跟着一起改了
	//切片可以基于一个数组灵活的创建新的数组
	fmt.Println(names1)
	fmt.Println(names)

	//切片完全独立于原始数组可以用copy函数
	//如果从0元素 开始截取，name冒号左边的数字 可以省略
	//字符串也可以进行切片截取
	sub1 := "hello world"[0:5]
	fmt.Println(sub1)

	//创建空切片的时候可以明确指定切片的容量，用make函数，可以提高运行效率
	str2 := make([]string, 10, 20) //指定长度和容量
	//names是定长数组，copy接受的是切片，使用[:]将数组变成切片
	copy(str2, names[:])
	fmt.Println(str2)
}
