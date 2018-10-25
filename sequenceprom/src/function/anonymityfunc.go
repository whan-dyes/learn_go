package function
import "fmt"

// abz:=func(a,b int,z float64) bool{
//   return (a*b)<int(z)
// }
func A(){
  v1 := func(i int) int {
    return i * i
  }
  fmt.Println(v1(3))

  v2:=b(2)
  fmt.Println(v2(3))

  for i:=0;i<3;i++{
    defer func(){
      fmt.Println("i=",i)
    }()
  }

  for i1:=0;i1<3;i1++{
     func(int){
      fmt.Println("i1=",i1)
    }(i1)
  }
}

func b(x int)func(int)int{
  return func(y int)int{
    return x+y
  }


}
