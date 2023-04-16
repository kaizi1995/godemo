package main

import (
	"fmt"
	"io/ioutil"
)

/*-------------条件判断demo----------*/
func bounded(v int) int {
	fmt.Println("----------基本条件判断----------")
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func one() {
	fmt.Println("---------第一种判断方式--------")
	const filename = "resource/abc.txt"
	file, err := ioutil.ReadFile(filename) //可以有两个返回值
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", file)
	}
	//fmt.Println(file)
}
func two() {
	fmt.Println("---------第一种判断方式简化版 --------")
	const filename = "resource/abc.txt"
	if file, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", file)
	}
	// fmt.Println(file) 这里都不能打印 file  因为生命周期不同（作用域不同）
}
func switchDemo(a, b int, op string) int {
	fmt.Println("---------switch demo  --------")
	var result int
	switch op { //不用break  会自动 break
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("no op :" + op) //panic 抛异常
	}
	return result
}

func grade(score int) string {

	fmt.Println("---------switch demo 没有表达式 --------")
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("取值范围不正确: %d", score))
	case score < 60:
		g = "D"
	case score < 70:
		g = "C"
	case score < 80:
		g = "B"
	case score < 90:
		g = "A"
	case score < 100:
		g = "S"
	}
	return g
}

func manyReturnDemo() (string, string, string) {
	fmt.Println("---------方法多个返回值 --------")
	return "kai zi", "NB", "666"

}
func fallthroughDemo(num int) {
	switch {
	case 1 == num:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 0 == num:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 1 == num:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 0 == num:
		fmt.Println("The integer was <= 7")
	case 1 == num:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}

func main() {

	fmt.Println(grade(4))
	fmt.Println(switchDemo(4, 2, "+"))
	fmt.Println(bounded(5))
	one()
	two()
	fmt.Println(manyReturnDemo())
	fallthroughDemo(1)
}

/*
知识点  go语言 定义都是反的

1.简化if条件语句 ，注意返回值的生命周期
2.switch case 里面不用break  绘自动 break 除非使用 fallthrough 参考：fallthroughDemo()
  switch 后面可以没有表达式
3. panic 抛异常
4.方法 是可以有多个返回值的


*/
