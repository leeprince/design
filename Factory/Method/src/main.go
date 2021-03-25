package main

import (
	"Method/src/caculate"
	"fmt"
)

func main() {
	add := caculate.AddResultFatory{}
	addFunc := add.Create()
	addFunc.SetLeft(10)
	addFunc.SetRight(20)
	fmt.Println("加法计算的结果：", addFunc.Result())

	m := caculate.MulResultFatory{}
	mulFunc := m.Create()
	mulFunc.SetLeft(10)
	mulFunc.SetRight(20)
	fmt.Println("乘法计算的结果：", mulFunc.Result())
}
