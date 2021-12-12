package main

import (
	"fmt"
	"os"
)

//defer 延迟可以修饰语句函数确保这条语句在栈退出时来执行
//一般做资源清理工作
//解锁关闭文件

func main() {
	filename := "01.Switch.go"
	readFile(filename)
}
func readFile(fileName string) {
	//一般go语言会将错误码作为最后一个参数返回
	f1, err := os.Open("01.Switch.go")
	//匿名函数 并立即调用
	defer func() {
		fmt.Println("准备关闭文件")
		_ = f1.Close()
	}()

	if err != nil {
		//出错了
		fmt.Println("打开文件失败 ", err)
	}
	buf := make([]byte, 1024)
	n, err := f1.Read(buf)
	fmt.Println("文件长度 ", n)
	fmt.Println("文件内容", string(buf))

	//f1.Close()

}
