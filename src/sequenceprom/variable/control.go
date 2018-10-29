package variable
import(
  . "fmt"
)
func Control(){
  Println("-------------------Control------------------")
  controltest1()
    controltest2()
}

func controltest1(){
  switch a:=1;{         //分号必须要，a 是 switch 局部变量
    case a>=0:Println("a=0")
    fallthrough
  case a>=1:
    Println("a=1")
  }

a1:=1
  switch a1{
    case 0:Println("a1=0")

  case 1:
    Println("a1=1")
  default:
    Println("None")
  }
}



func controltest2(){
LABEL1:
for{
  for i:=0;i<10;i++{
    if i>3{
        break LABEL1    //goto  LABEL1 会成无限循环（LABEL1只能放在后面），只有 break 只能跳出一层循环
    }                   //continue LABEL1 会跳到 LABEL1 ,也会无限循环

  }
  Println("OK")
}
}
func controltest3(){

}
func controltest4(){

}
func controltest5(){

}
func controltest6(){

}
func controltest7(){

}
/*


*/
