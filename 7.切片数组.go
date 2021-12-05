package main

import "fmt"

func main() {
	//切片slice  他的底层也是数组，但是可以动态改变长度
	//定义一个切片,包含多个地名
	names := []string{"北京", "上海", "广州", "深圳"}
	for i, v := range names {
		fmt.Println("i is ", i, "v is ", v)
	}

	names = append(names, "海口", "石家庄") //追加 赋值回原来的切片
	for i, v := range names {
		fmt.Println("i is ", i, "v is ", v)
	}

	//2.对于一个切片，不仅有长度的概念len（），还有一个容量的概念cap（）
	fmt.Println(len(names))
	fmt.Println(cap(names))

}
