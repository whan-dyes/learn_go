package variable
import "fmt"

func BoolVar(){
  var v1,v2 bool
  v1,v2=true,false
v3 :=(1==2)  // v3也会被推导为bool类型
// 布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换。以下的示例是一些错误
// 的用法，会导致编译错误：
// var b bool
// b = 1 // 编译错误
// b = bool(1) // 编译错误

  fmt.Println("bool_v1=",v1,"bool_v2=",v2,"v3=",v3)
}
