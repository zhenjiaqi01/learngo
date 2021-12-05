package main

import "fmt"

func main() {
	//1-定义

	//nums:=[10]int{1,2,3,4}//常用方式
	//var nams = [10]int{1,2,3,4}
	//var nums[10]int = [10]int{1,2,3,4}
	nums := [10]int{1, 2, 3, 4}
	//遍历 方式一
	for i := 0; i < len(nums); i++ {
		fmt.Println("i:", i, ",j: ", nums[i])
	}

	//遍历方式二 for range
	for key, value := range nums {
		//value是临时变量，不断地被重新赋值，修改它不会修改原始数组
		//golang中 ：= 左边必须有新变量
		value += 1
		fmt.Println("key:", key, "value:", value, "nums", nums[key])
	}

	//3-使用make进行创建数组（不定长数组）
}
