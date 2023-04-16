package main

import "fmt"

func simpleDemo() {
	// 基本数据类型在内存中的布局
	var i int = 100
	// i 的地址是什么
	fmt.Println("i 的地址是 ", &i)

	//指针

	//ptr 是一个指针变量 类型是 *int  本身的值是 &i(i在内存中的地址)
	//不能将一个值给指针,指针接收地址
	var ptr *int = &i //var ptr *int =i  (×)
	fmt.Printf("ptr 的值是%v\n", ptr)
	//ptr 的地址
	fmt.Printf("ptr 的地址是%v\n", &ptr)
	//获取指针类型所指向的值
	fmt.Printf("ptr指针 指向的值%v\n", *ptr)

	i = 99
	fmt.Printf("&i 的地址是%v\n", &i)

	fmt.Printf("i 的值是%v\n", i)

}

func simpleDemo1() {
	fmt.Println("------------------------")

	//获取一个变量的地址
	var i1 int = 99
	fmt.Println("i1 的地址是", &i1)
	var ptr1 *int = &i1
	fmt.Printf("ptr1的值是%v\n", ptr1)
	*ptr1 = 98
	fmt.Printf("修改后ptr1指向的值是%v\n", *ptr1)
}

func main() {
	simpleDemo()
	simpleDemo1()
}

/*注意点
1。 指针不能计算
2。 指针只能接收地址 不能接收 值



*/
