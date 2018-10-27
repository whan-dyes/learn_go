package variable
import "fmt"
import "math"
// 全局变量申明，只能用在函数体外
var(
  name int=100
  Name int=200
)
func IsEqual(f1, f2, p float64) bool {
return math.Dim(f1, f2) < p
}
func FloatVar(){

  // 浮点型用于表示包含小数点的数据，比如1.234就是一个浮点型数据。Go语言中的浮点类型
  // 采用IEEE-754标准的表达方式。
  // 1. 浮点数表示
  // Go语言定义了两个类型float32和float64，其中float32等价于C语言的float类型，
  // float64等价于C语言的double类型。
  // 在Go语言里，定义一个浮点数变量的代码如下：
  var fvalue1 float32
  fvalue1 = 12
  fvalue2 := 12.0 // 如果不加小数点，fvalue2会被推导为整型而不是浮点型
  // 对于以上例子中类型被自动推导的fvalue2，需要注意的是其类型将被自动设为float64，
  // 而不管赋给它的数字是否是用32位长度表示的。因此，对于以上的例子，下面的赋值将导致编译
  // 错误：
  // fvalue1 = fvalue2
  // 而必须使用这样的强制类型转换：
  fvalue1 = float32(fvalue2)
  // 2. 浮点数比较
  // 因为浮点数不是一种精确的表达方式，所以像整型那样直接用==来判断两个浮点数是否相等
  // 是不可行的，这可能会导致不稳定的结果。
  // 下面是一种推荐的替代方案：

  // p为用户自定义的比较精度，比如0.00001
  var f1,f2 float64=1.334,1.2340
  var p float64=0.0001

if IsEqual(f1,f2,p)==true{
fmt.Println("f1==f2")
}else{
fmt.Println("f1!=f2")
}


  fmt.Println("bool_v1=",fvalue1,"bool_v2=",fvalue2)
}
