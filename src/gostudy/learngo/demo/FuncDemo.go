package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

func simple(a, b int) int {
	fmt.Print("-------普通方法------")
	return a + b
}

func manyReturn(a, b int) (int, int, string) {
	fmt.Print("-------多个返回值------")
	return 0, 1, "hahaha"
}

func assignReturn() (q, r int) {
	fmt.Print("-------指定返回值------")
	q = 1
	r = 2
	return
	// 或者 return 1,2
}

func receiveOneReturn() int {
	fmt.Print("-------多个返回值 只需要一个------")
	q, _ := assignReturn() //可以用占位符 _ 代替
	return q
}

func compatibility(a, b int, op string) (int, error) {
	fmt.Println("---------容错处理  --------")
	switch op { //不用break  绘自动 break
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("操作类型不存在")
	}

}

func funcParam(op func(int, int) int, a, b int) int {
	fmt.Println("---------参数是方法的方法  --------")
	pointer := reflect.ValueOf(op).Pointer()    //通过reflect 代理 获取指针 Pointer()
	opName := runtime.FuncForPC(pointer).Name() //获取运行时方法名称
	fmt.Printf("获取的方法名称 %s 和参数"+"(%d,%d)\n", opName, a, b)
	return op(a, b)

}

func sumDemo(nums ...int) int {
	fmt.Println("---------可变参数列表  --------")
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(simple(1, 0))
	a, b, c := manyReturn(1, 0)
	fmt.Println(a, b, c)
	fmt.Println(assignReturn())
	fmt.Print(receiveOneReturn())
	if rep, err := compatibility(4, 2, "l"); err == nil {
		fmt.Println(rep)
	} else {
		fmt.Println(err)
	}
	fmt.Println(funcParam(simple, 1, 3))

	fmt.Printf("匿名函数 返回值%s 测试", strconv.Itoa(funcParam(func(a int, b int) int {
		return a + b
	}, 1, 3)))

	fmt.Print(sumDemo(1, 2, 3))

}

/*
	知识点  返回值类型在后面   先名称后类型
	1.  可以多个返回值
	2。可以指定返回值
	3。 多个返回值值只需要一个  可以用_ 替代
	4.方法的参数可以是方法  参考 funcParam（）
	5.	pointer := reflect.ValueOf(op).Pointer()    //通过reflect 代理 获取指针 Pointer()
	6. runtime.FuncForPC(pointer).Name() //获取运行时方法名称
	7.range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
		for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环

*/
