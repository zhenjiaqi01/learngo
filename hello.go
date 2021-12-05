package main //每个go程序都必须有package main
//每个go程序都是.go结尾的

//导入一个标准包fmt，format，format一般用于格式化输入
import "fmt"

//主函数，所有函数都必须func开头
//一个 函数返回值不是放在函数前面，而是放在后面
func main() {
	//go语言不需要分号结尾。
	fmt.Println("Hello world")
}
