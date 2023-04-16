package main

import "fmt"

func arrDemo() {
	var arr1 [3]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 3, 5, 6}
	fmt.Print(arr1, arr2, arr3)
	var grid [4][5]int
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i]) //for循环遍历
	}

	for i := range arr3 {
		fmt.Println(i) //range循环遍历 下标
	}

	for _, v := range arr3 {
		fmt.Println(v) //range循环遍历 值
	}

	for i, v := range arr3 {
		fmt.Println(i, v) //range循环遍历 下标和 值
	}

}

func printArray(arr [3]int) {
	arr[0] = 100 //数组是值类型传递
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray1(arr *[3]int) {
	arr[0] = 100 //数组是值类型传递 ，如果是指针类型的话 可以打破值传递
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	//arrDemo()
	arr1 := [...]int{1, 2, 3}
	printArray(arr1)
	fmt.Println(arr1)

	printArray1(&arr1)
	fmt.Println(arr1)

}

/*
注意点
1。 :=需要指定值
2。 [...] 可以动态扩容
3。 数组长度在前  类型在后
4.range 可以遍历  且 返回值有下标和值
5。range 可以用 _代替返回值 ，只取值
6。 数组是值类型传递（拷贝）
7。go 中一般不直接使用数组 ，会使用切片



*/
