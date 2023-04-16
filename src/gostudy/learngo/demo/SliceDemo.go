package main

import "fmt"

//SimpleDemo 基本语法
func SimpleDemo() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = arr[2:6] //切片
	fmt.Printf("切片基本语法arr[2:6]%d\n", s)
	fmt.Printf("切片基本语法arr[:6]%d\n  ", arr[:6])

	fmt.Printf("切片基本语法[2:]%d\n", arr[2:])
	s1 := arr[2:]
	updateSlice(s1)
	fmt.Printf("切片基本语法[2:]当作参数传入修改后%d\n", arr[2:])

	fmt.Printf("切片基本语法arr[:]%d\n", arr[:])
	s2 := arr[:]
	updateSlice(s2)
	fmt.Printf("切片基本语法arr当作参数传入修改后[:]%d\n", arr[:])

}

func updateSlice(arr []int) { //当作参数传递时
	arr[0] = 100
}

func resliceDemo() { //   数组转切片
	var s = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := s[1:6]
	fmt.Println(s1)
	s1 = s1[2:]
	fmt.Println(s1)
	fmt.Println(s)
}

func sliceExtendDemo() { //  扩展值 （隐藏值）
	var s = [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	s1 := s[1:6] //s1 对应下标值是  0-b ,1-c ,2-d,3-e,4-f ,g,h 其中 是扩展值 （隐藏值）
	fmt.Println("s1=s[1:6]=", s1, ",len(sl) =", len(s1), ",cap(sl) =", cap(s1))
	s2 := s1[3:5] //s2 对应下标值是  0-e ,1-f ,g,h 其中  g，h 是扩展值 （隐藏值）
	fmt.Println("s2 = s1[3:5]=", s2, ",len(s2) =", len(s2), ",cap(s2) =", cap(s2))
	s3 := s2[1:4] //s3 对应下标值是  0-f ,1-g,2-h 其中 1-g,2-h 是原数组的隐藏值 被使用了
	fmt.Println("s3 = s2[1:4]=", s3, ",len(s3) =", len(s3), ",cap(s2) =", cap(s2))
	fmt.Println(s)
}

//增加元素
func appendDemo() {
	var s = [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	s1 := s[3:6]
	s1 = append(s1, "i") //append 可以一个一个元素添加
	s2 := append(s1, "j")
	s3 := append(s2, "k", "l") //可以一个一个元素迭起添加
	s4 := append(s3, s1[:]...) //可以添加一个切片
	fmt.Println(s)
	fmt.Println("s1 = ", s1, " s2 = ", s2, " s3 = ", s3, " s4 = ", s4)

}

func BaseDemo() {
	var s []int // 可以直接定义切片   零值是nil
	for i := 0; i < 100; i++ {
		fmt.Println(" len(s)= ", len(s), " cap(s)= ", cap(s))
		s = append(s, i)
	}
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4}   //定义默认值的切片
	s2 := make([]int, 15)     //定义定长的切片
	s3 := make([]int, 10, 32) ////定义定长 长度10 cap32的切片
	fmt.Println(" len(s1)= ", len(s1), " cap(s1)= ", cap(s1))
	fmt.Println(s2, " len(s2)= ", len(s2), " cap(s2)= ", cap(s2))
	fmt.Println(s3, " len(s3)= ", len(s3), " cap(s3)= ", cap(s3))
	copy(s2, s1) // s1 把值 copy给s2
	fmt.Println(s2, " len(s2)= ", len(s2), " cap(s2)= ", cap(s2))

	s2 = append(s2[:3], s2[4:]...) //删除元素
	fmt.Println(s2, " len(s2)= ", len(s2), " cap(s2)= ", cap(s2))

}

func main() {
	//SimpleDemo()
	//resliceDemo()
	//sliceExtendDemo()
	//appendDemo()
	BaseDemo()

}

/*
知识点
 1。 slice （切片）可以看作数组的视图 [1:4]  下标【1，6）
 2，slice （切片） 数组型 是引用传递
 3， 可以对当前切片 连续处理 resliceDemo()
 4。在原有数组上面进行切片处理时 可以往后扩展/取隐藏值获取（但是不能超过cap 长度） 但是不可以往前扩展/取隐藏值 具体看:sliceExtendDemo()
  ，可以参考图片 slice 和sliceExtend
 5。 slice 三个关键字  prt 指向 slice 开头的元素  ,len 指slice长度是多少,取值 只能在这个len 范围内 否则下标越界
	cap 代表 整个 ptr 开始到结束到长度
	s[i] 不能超过 len（s)的长度 ，往后扩展不能超过底层数组cap(s) 长度

6. append 可以一个一个元素添加(append(s,1,2,3)) 也可以是切片 （append(s,a[3:])）， 对原有的数组切片 然后添加新值 ，会填充到固定到cap长度，
    如果超越 cap 元素  系统会重新分配更大的底层数组c,原有数组会被垃圾回收掉（没人用到情况下）  demo :appendDemo()
7。 由于 值传递到关系，必须接收append的返回值   s= append(s,val)
8. 直接定义切片 s := []int  零值是nil
9. cap 的扩容  是2n ，当值是nil 时候 cap是0，1，2，4.....  具体看 BaseDemo（）
10.copy(s2, s1)  // s1 把值 copy给s2
11. delete删除元素 s2 = append(s2[:3], s2[:4]...)


*/
