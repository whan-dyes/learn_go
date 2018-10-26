package variable
import(
  "fmt"
)
  func ComplexVar(){
    var value1 complex64 // 由2个float32构成的复数类型
    value1 = 3.2 + 12i
    value2 := 3.2 + 12i // value2是complex128类型
    value3 := complex(3.2, 12) // value3结果同 value2
    fmt.Println("value1=",value1)
    fmt.Println("value2=",value2)
    fmt.Println("value3=",value3)

    // 2. 实部与虚部
    // 对于一个复数z = complex(x, y)，就可以通过Go语言内置函数real(z)获得该复数的实
    // 部，也就是x，通过imag(z)获得该复数的虚部，也就是y。
    // 更多关于复数的函数，请查阅math/cmplx标准库的文档。
    fmt.Println("value2=",real(value2),"+",imag(value2),"i")
  }
