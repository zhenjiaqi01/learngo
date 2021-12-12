package main

import "fmt"

func main() {
	//标签label1
	//goto label1
	//break label1
	//continue label1
LABEL:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//goto LABEL
				//continue LABEL
				break LABEL
			}
			fmt.Println("i: ", i, "j: ", j)
		}
	}
}
