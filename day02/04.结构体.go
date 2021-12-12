package main

import "fmt"

//c:
//struct Person{}

//c语言里有typedef int myInt

//go语言结构体 type + struct来处理
type Person struct {
	name   string
	age    int
	gender string
	score  float64
}

type MyInt int

func main() {
	var i, j MyInt
	i, j = 10, 20
	fmt.Println("i+j is ", i+j)

	zhen := Person{
		name:   "zhen jiaqi",
		age:    0,
		gender: "female",
		score:  90,
	}

	fmt.Println("zhen", zhen.name, " age ", zhen.age, "blah blah")
	//指针
	s1 := &zhen
	fmt.Println((*s1).gender, s1.score)
}
