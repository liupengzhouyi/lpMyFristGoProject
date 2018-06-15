package main

import ."fmt"
import "unsafe"
//为包起 别名
//import stream "io"

var (
	// 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b int
)

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
	var v_name string = "value"
	var v_fales bool
	Println(v_name)
	Println("esxrdcftvg", v_fales)


	var v_name1, v_name2,v_name3 int
	v_name1, v_name2, v_name3 = 1,2,3
	Println(v_name1,v_name2,v_name3)



	//交换俩个变量的值
	a = 12
	b = 21
	Println(a, b)
	a, b = b, a
	Println(a, b)


	//常量
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH
	Println("面积：", area)

	var lp_string = "abcdefg"
	var lp_number = len(lp_string)
	Println("长度：", lp_number)
	Println(unsafe.Sizeof(lp_string))



 }























