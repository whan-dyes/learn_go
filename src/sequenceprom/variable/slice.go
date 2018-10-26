package variable

import(
  "fmt"
)

func Slice(){
  // 先定义一个数组
  var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  // 基于数组创建一个数组切片
  var mySlice []int = myArray[:5]    //前面5个元素
  fmt.Print("Elements of myArray: ")
  for _, v := range myArray {
    fmt.Print(v, " ")
  }
  fmt.Print("\nElements of mySlice: ")
  for _, v := range mySlice {
    fmt.Print(v, " ")
  }
  fmt.Println()

  var mySlice1 []int = myArray[5:] //后面5个元素
  fmt.Print("mySlice1=")
  for _, v := range mySlice1 {
    fmt.Print(v, " ")
  }
  fmt.Println()

  fmt.Print("mySlice2=")
  var mySlice2 []int = myArray[:] //所有元素
  for _, v := range mySlice2 {
    fmt.Print(v, " ")
  }
  fmt.Println()

  var mySlice3 []int =make([]int,5,10) //所有元素
  fmt.Print("mySlice3=")
  for _, v := range mySlice3 {
    fmt.Print(v, " ")
  }
  fmt.Println()

  fmt.Print("mySlice4=")
  mySlice4:=append(mySlice3,6,7,8)
  fmt.Print(mySlice4)
  fmt.Println()

  fmt.Print("mySlice5=")
  mySlice5:=append(mySlice3,mySlice1...)
  fmt.Println(mySlice5)
  fmt.Println("mySlice5.len()=",len(mySlice5))
  fmt.Println("mySlice5.cap()=",cap(mySlice5))
  fmt.Println()

  slice1 := []int{1, 2, 3, 4, 5}
  slice2 := []int{5, 4, 3}
  copy(slice1, slice2) // 只会复制slice1的前3个元素到slice2中
  fmt.Println("slice1=",slice1)
  fmt.Println("slice2=",slice2)
  copy(slice2, slice1[2:]) // 只会复制slice2的3个元素到slice1的前3个位置
  fmt.Println("slice1=",slice1)
  fmt.Println("slice2=",slice2)

  s1:=[]int{1,2,3,4,5}
  s2:=s1
  fmt.Println("s1=",s1)
  fmt.Println("S2=",s2)

  m1:=map[int]string{1:"a",2:"b",3:"c"}
  fmt.Println("m1=",m1)
  m2:=make(map[string]int)

  for k,v:=range m1{
    m2[v]=k
  }
  fmt.Println("m2=",m2)
}
