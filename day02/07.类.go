package main

import "fmt"

//Person类，绑定方法
type Person01 struct {
	name   string
	age    int
	gender string
	score  float64
}

//类外面绑定方法
//func(p Person01)Eat(){
//	fmt.Println("Person is eatting")
//}
//第二种方法绑定可以用指针
func (this *Person01) Eat() {
	fmt.Println("人吃饭")
}
func main() {
	lily := Person01{
		name:   "Lily",
		age:    30,
		gender: "女",
		score:  0,
	}
	fmt.Println(lily)
	lily.Eat()
}
