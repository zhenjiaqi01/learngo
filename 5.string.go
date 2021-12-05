package main

import "fmt"

func main() {
	name := "Duke"
	//需要换行用反引号比如
	usage := `./a.out <option>
			-h help
			-a xxxxxx
			`
	fmt.Println("name : ", name)
	fmt.Println("usage ", usage)

	//访问长度
	l1 := len(usage)
	fmt.Println(l1)

	//不需要加括号
	for i := 0; i < len(name); i++ {
		fmt.Printf("i: %d,v: %c\n", i, name[i])
	}

	//字符串拼接
	i, j := "hello", "world"
	fmt.Println("i+j = ", i+" "+j)

	//常量字符串
	const address = "Beijing" //常量就不能自动推导了
	fmt.Println(address)

}
