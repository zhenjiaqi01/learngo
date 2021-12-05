package main

import (
	"fmt"
	"os"
)

//从命令行输入参数，switch中进行判断
func main() {
	//os包下的Args 直接获取命令输入，是一个字符串切片
	cmds := os.Args
	//os.Args[0] 是程序名字
	//os.Args[1]是第一个参数
	for key, cmd := range cmds {
		fmt.Println("key ", key, "cmd ", cmd)
	}
	if len(cmds) < 2 {
		fmt.Println("请正确输入参数")
	}
	switch cmds[1] {
	//go 中的switch默认帮你加好了break
	//如果想向下穿透可以使用fallthrough
	case "hello":
		fmt.Println("Hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default called")
	}
}
