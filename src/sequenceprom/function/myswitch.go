package function
import(
  "fmt"
)

func MySwitch(){
  i:=2
  switch i {
  case 0:
    fmt.Printf("0")
  case 1:
    fmt.Printf("1")
    case 2:     // i = 2时，输出3；
    fallthrough
    case 3:     // i = 3时，输出3；
    fmt.Printf("3")
  case 4, 5, 6:
    fmt.Printf("4, 5, 6")
  default:
    fmt.Printf("Default")
}
fmt.Println()

num:=8
    switch {
    case 0 <= num && num <= 3:
      fmt.Printf("0-3")
    case 4 <= num && num <= 6:
      fmt.Printf("4-6")
    case 7 <= num && num <= 9:
      fmt.Printf("7-9")
    }

  // 在使用switch结构时，我们需要注意以下几点：
  //  左花括号{必须与switch处于同一行；
  //  条件表达式不限制为常量或者整数；
  //  单个case中，可以出现多个结果选项；
  //  与C语言等规则相反，Go语言不需要用break来明确退出一个case；
  //  只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；
  //  可以不设定switch之后的条件表
}
