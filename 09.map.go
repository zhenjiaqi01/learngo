package main

import "fmt"

func main() {
	//定义一个字典
	//var idNames map[int]string
	////make 手动分配空间 可以不指定长度
	//idNames = make(map[int]string,10)
	//idNames[0] = "Duke"
	//idNames[1] = "Lily"

	//定义直接赋值
	idNames := make(map[int]string, 10)
	for key, value := range idNames {
		fmt.Println("key ", key, "value", value)
	}

	//如何确定一个key是否存在map中
	//map不存在访问越界的问题 他认为所有key都是有效的，所以访问一个不存在的key不会崩溃，会返回这个类型的零值
	name9 := idNames[111111]
	fmt.Println(name9)

	//校验key是否存在的方式
	zhi, cunzai := idNames[1]
	fmt.Println(cunzai)
	if cunzai {
		fmt.Println(zhi)
	} else {
		fmt.Println("不存在噻")
	}

	//删除map中元素
	delete(idNames, 1)

	//多并发处理需要对map上锁

}
