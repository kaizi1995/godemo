package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func forEachDemo() {
	fmt.Println("----------for循环小demo--")
	sum := 0
	for i := 1; i < 100; i++ {
		sum++
	}
	fmt.Printf("%s\n", strconv.Itoa(sum))
}

/**/
func convertToBin(n int) { //for的条件里可以省略初始条件，结束条件 ，递增表达式
	fmt.Printf("%s\n", "----------转换二进制--")
	result := ""
	if n == 0 {
		result = "0"
	} else {
		for ; n > 0; n /= 2 {
			lsb := n % 2
			result = strconv.Itoa(lsb) + result
		}
	}
	fmt.Println(result)
}
func printFile(filename string) {
	fmt.Println("----------一行一行输出文件内容--")
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	fmt.Println("----------死循环--")
	for {
		fmt.Println("死循环")
	}
}

func main() {
	forEachDemo()
	convertToBin(5)
	convertToBin(13)
	convertToBin(0)
	printFile("resource/abc.txt")

}

/*
	知识点
	1.for 不用括号 和if 一样都不需要括号
	2.for的条件里可以省略初始条件，结束条件 ，递增表达式
	3. for {} 死循环 go 里面没有while
	4. strconv.Itoa() 数字转字符串



*/
