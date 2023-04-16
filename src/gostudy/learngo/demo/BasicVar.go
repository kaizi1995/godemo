package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*-------------变量，常量，枚举的定义-----------*/

var aa = "NB" //包内变量（作用于当前包） 不能使用 : 去定义
var bb = 666

var ( //多个var  可以用（）去概括
	cc = "hahaha"
	dd = 888
)

func initVar() {
	fmt.Println("-------基本类型参数定义----------")
	var num int = 3
	var a, b int = 0, 1
	var str string = "king"
	fmt.Println(num, str, a, b)
}

func initVarType() {
	fmt.Println("-------var多参数自适配类型 参数定义----------")
	var a, b, c, d = 1, true, "yes", 5 //不定义类型 可以多个参数拼接一起 自适配类型
	fmt.Println(a, b, c, d)
}
func initVarkuohao() {
	fmt.Println("-------var多参数括号自适配类型 参数定义----------")
	var (
		a, b, c, d = 6, false, "kuohao", 6
	) //括号

	fmt.Println(a, b, c, d)
}

func initVarShortType() {
	fmt.Println("------- ： 替代 var 参数定义----------")
	a, b, c, d := 1, true, "yes", 5 //: 等价于  var
	d = 2
	fmt.Println(a, b, c, d)
}

func variable() {
	fmt.Println("------- 空字符串输出 参数定义----------")
	var num int
	var str string
	fmt.Printf("%d %q\n\n", num, str)
	fmt.Println("") //空字符串 不能直接写 ""

}

func complexDemo() {
	fmt.Println("-------复数----------")
	var score complex64 = complex(1, 2)
	var number complex128 = complex(23.23, 11.11)
	fmt.Print("Score = ", score, " Number = ", number)
}

func euler() {
	fmt.Println("\n\n", "-------欧拉公式----------")
	a := 3 + 4i
	b := cmplx.Pow(math.E, 1i*math.Pi) + 1
	c := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(cmplx.Abs(a), b, c)
	fmt.Printf("%.3f\n", c)
}

func constVar() {
	fmt.Println("-------常量----------")
	const name string = "av.txt"
	const name1 = "ok"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(name, name1, a, b, c)

}

func enums() {
	fmt.Println("-------枚举----------")
	const (
		java   = 1
		python = 2
		golang = 5
	)

	const (
		zero = iota //iota 自增
		one
		_ //   _ 相当于占位符
		three
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	const (
		a = 5 //  a=5  c ，d 都会赋值 5
		c
		d
	)

	fmt.Println(java, python, golang)
	fmt.Println(zero, one, three)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(a, c, d)

}

func main() {
	/*	fmt.Println("hello kai zi ", aa, bb, cc, dd)
		variable()
		initVar()
		initVarType()
		initVarShortType()
		initVarkuohao()
		complexDemo()
		euler()
		constVar()*/
	enums()

}

//知识点  变量类型  （ 先名称后类型）
/*
 1.bool
 2.string
 3.(u)int, (u)int8, (u)int16, (u)int32 ,(u)int64, uintptr(指针)
  注意点： 加u 无符号整数(整数 取值范围  2^8 -1 = 255),  不加u有符号整数（-2^7 = -128 ~ 2^7-1 = 127） ,不加数字是根据操作系统来设置默认大小
 4.byte “byte是uint8的内置别名” 等于 8比特（bit）
 5.rune（类似char） 字符型  长度32位   “rune是int32的内置别名”
 6.float32 ,float64 ,complex64 ,complex128
   注意点： complex(复数) complex64中实部和虚部是32位的，在complex128中实部和虚部是64位的。
	参考：complexDemo()
 7. const 常量，  iota 自增数  ,定义自增枚举时候 _ 相当于一个占位符一样
	如果常量第一个值赋值，后面的定义了但是没有赋值的话都会默认使用第一个值
	参考：具体看enums()
*/
