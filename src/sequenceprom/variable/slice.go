package variable

import (
	. "fmt"
)

/*
切片Slice

其本身并不是数组，它指向底层的数组
作为变长数组的替代方案，可以关联底层数组的局部或全部
为引用类型
可以直接创建或从底层数组获取生成
使用len()获取元素个数，cap()获取容量
一般使用make()创建
如果多个slice指向相同底层数组，其中一个的值改变会影响全部

make([]T, len, cap)
其中cap可以省略，则和len的值相同
len表示存数的元素个数，cap表示容量

Reslice时索引以被slice的切片为准
索引不可以超过被slice的切片的容量cap()值
索引越界不会导致底层数组的重新分配而是引发错误

Append
可以在slice尾部追加元素
可以将一个slice追加在另一个slice尾部
如果最终长度未超过追加到slice的容量则返回原始slice
如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据

Copy

*/

func Slice() {
	Println("---------------------Slice-----------------------------------------")
	slice1()
	slice2()
	slice3()
	slice4()
	slice5()
	slice6()
	slice7()
	slice8()
}
func slice1() {
	Println(".........slice1.........")
	// 先定义一个数组
	// var myArray  = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var myArray = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5] //前面5个元素
	Print("Elements of myArray: ")
	for _, v := range myArray {
		Print(v, " ")
	}
	Print("\nElements of mySlice: ")
	for _, v := range mySlice {
		Print(v, " ")
	}
	Println()

	var mySlice1 []int = myArray[5:10] //从第 5 位后面的元素开始至第 10 个元素
	Print("mySlice1=")
	for _, v := range mySlice1 {
		Print(v, " ")
	}
	Println()

	Print("mySlice2=")
	var mySlice2 []int = myArray[:] //所有元素
	for _, v := range mySlice2 {
		Print(v, " ")
	}
	Println()

	var mySlice3 []int = make([]int, 5, 10)
	Print("mySlice3=")
	for _, v := range mySlice3 {
		Print(v, " ")
	}
	Println()

	Print("mySlice4=")
	mySlice4 := append(mySlice3, 6, 7, 8)
	Print(mySlice4)
	Println()

	Print("mySlice5=")
	mySlice5 := append(mySlice3, mySlice1...)
	Println(mySlice5)
	Println("mySlice5.len()=", len(mySlice5))
	Println("mySlice5.cap()=", cap(mySlice5))
	Println()

}
func slice2() {
	Println(".........slice2.........")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice1, slice2) // 只会复制slice1的前3个元素到slice2中
	Println("slice1=", slice1)
	Println("slice2=", slice2)
	copy(slice2, slice1[2:]) // 只会复制slice2的3个元素到slice1的前3个位置
	Println("slice1=", slice1)
	Println("slice2=", slice2)

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	Println("s1=", s1)
	Println("S2=", s2)

	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	Println("m1=", m1)
	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k
	}
	Println("m2=", m2)
}
func slice3() {
	Println(".........slice3.........")
	/*slice指向底层数组，其中1个改变，另1个随着改变，当其中1个append超界后，指向
	  了另 1 个重新分配的数组，他就不随着改变了。
	*/
	s1 := make([]int, 3, 6)
	Printf("%v,%p \n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	Printf("%v %p\n", s1, s1)

	a := []int{1, 2, 3, 4, 5}
	s2 := a[2:5]
	s3 := a[1:3]
	Printf("%v,%p \n", s2, s2)
	Printf("%v,%p \n", s3, s3)

	s3[1] = 100
	Println(s2, "\n", s3, "\n", a[2])
}
func slice4() {
	Println(".........slice4.........")
	var sm = make([]map[int]string, 5, 10)
	for i, _ := range sm {
		sm[i] = make(map[int]string)
		for j := 0; j < 2; j++ {
			sm[i][j] = "good"
		}
	}
	Println(sm)

	for j := 0; j < 5; j++ {
		sm[j] = make(map[int]string)
		for k := 0; k < 3; k++ {
			sm[j][k] = "ok"
		}
	}

	Println(sm)
}
func slice5() {
	Println(".........slice5.........")
}
func slice6() {
	Println(".........slice6.........")
}
func slice7() {
	Println(".........slice7.........")
}
func slice8() {
	Println(".........slice8.........")
}
func slice9() {
	Println(".........slice9.........")
}
