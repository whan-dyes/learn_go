package function
import (
  . "fmt"
)

// abz:=func(a,b int,z float64) bool{
//   return (a*b)<int(z)
// }
func AnonymityFun(){
  Println("<<<<<<<<<<<<<<<<<<<<<AnonymityFun>>>>>>>>>>>>>>>>>>>>>>>>")
  anonymity1()
}

func anonymity1(){
  Println("................anonymity1................")
  v1 := func(i int) int {
    return i * i
  }
  Println(v1(3))

  v2:=b(2)
  Println(v2(3))

  for i:=0;i<5;i++{
    defer func(){       //defer 在函数返回后才运行，这时 i 已经变成了 5
      Println("i=",i)
      }()
    }

    for i1:=0;i1<3;i1++{
      func(i int){
        Println("i1=",i)
        }(i1)
      }
    }
    func b(x int)func(int)int{
      return func(y int)int{
        return x+y
      }


    }
