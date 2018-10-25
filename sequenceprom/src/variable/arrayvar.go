package variable
import(
  "fmt"
)

func PrintArray(arr [5]int){
  for i,n:=range arr{
    fmt.Println("arr[",i,"]=",n)
  }
  fmt.Println("arr=",arr)
  fmt.Println()
}

func ArrayVar(){
  var arr [5]int=[5]int{1,2,3,4,5}

  for i:=0;i<len(arr);i++{
    n:=arr[i]
    fmt.Println("arr[",i,"]=",n)
  }
  fmt.Println()

  PrintArray(arr)

  a:=[20]int{19: 1}
  b:=[...]int{1,2,3,4,5}
  c:=[...]int{19: 2}

  var pb *[5]int=&b  //指向数组指针 pb= &[1 2 3 4 5]
  x,y:=1,2
  var d =[...]*int{&x,&y}  //保存*int的数组 d= [0xc042054408 0xc042054410]

fmt.Println("a=",a)
fmt.Println("b=",b)
fmt.Println("c=",c)
fmt.Println("d=",d)
fmt.Println("pb=",pb)
fmt.Println("a==c",a==c)
fmt.Println("b==*pb",b==*pb)
fmt.Println("&b==pb",&b==pb)
fmt.Println()

i:=[2][3]int{ //行数可以自动计算  i:=[...][3]int 列不行
  {0:1,2,3},
  {10,20,2:30}}
fmt.Println("i=",i)

}
// 以下为一些常规的数组声明方法：
// [32]byte // 长度为32的数组，每个元素为一个字节
// [2*N] struct { x, y int32 } // 复杂类型数组
// [1000]*float64 // 指针数组
// [3][5]int // 二维数组
// [2][2][2]float64 // 等同于[2]([2]([2]float64))
// 从以上类型也可以看出，数组可以是多维的，比如[3][5]int就表达了一个3行5列的二维整
// 型数组，总共可以存放15个整型元素。

// 在Go语言中，数组长度在定义后就不可更改，在声明时长度可以为一个常量或者一个常量
// 表达式（常量表达式是指在编译期即可计算结果的表达式）。数组的长度是该数组类型的一个内
// 置常量，可以用Go语言的内置函数len()来获取。下面是一个获取数组arr元素个数的写法：
// arrLength := len(arr)


//第一种
//var <数组名称> [<数组长度>]<数组元素>
// var arr [2]int
//     arr[0]=1
//     arr[1]=2
//
// //第二种
// //var <数组名称> = [<数组长度>]<数组元素>{元素1,元素2,...}
// var arr = [2]int{1,2}
// //或者
// arr := [2]int{1,2}
//
// //第三种
// //var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{元素1,元素2,...}
// var arr = [...]int{1,2}
// //或者
// arr := [...]int{1,2}
//
// //第四种
// //var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{索引1:元素1,索引2:元素2,...}
// var arr = [...]int{1:1,0:2}
// //或者
// arr := [...]int{1:1,0:2}
