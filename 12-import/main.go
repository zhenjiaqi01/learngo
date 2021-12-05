package main

import (
	"12-import/add"
	"fmt"
)

//go语言中，同一层级目录不允许出现多个包名
import (
	SUB "12-import/sub" //sub是文件名也是包名 顺便改名叫SUB
	//."12-import/sub"前面加个点 这样在下面用的时候就不用包名点函数名了 而是可以直接用函数名
)

//sub是文件名，同时也是包名
func main() {
	res := SUB.Sub(2, 1) //包名.函数去调用
	fmt.Println(res)

	//如果包里的函数想要被外部访问，name一定要手写字母大写，如果是小写开头 那只有同报名的文件
	add.Add(1, 1)
}
