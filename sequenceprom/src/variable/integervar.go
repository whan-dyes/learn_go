package variable
import "fmt"

func IntegerVar(){
  var v10,v11 int=1,3
  var v12=100
  v13:=130
  var v14 int
  v14=140
  // v14:=140     no new variables on left side of :=

  var (
    v90 int=9000
    v91 int=9100
  )
  fmt.Println("v10=",v10,"v11=",v11,"v12=",v12,"v13=",v13,"v14=",v14,"v90=",v90,"v91=",v91)
}
// 类 型              长度（字节）                      值 范 围
// int8                  1                           -128 ~ 127
// uint8（即byte）        1                           0 ~ 255
// int16                  2                         -32 768 ~ 32 767
// uint16                 2                           0 ~ 65 535
// int32                  4                   -2 147 483 648 ~ 2 147 483 647
// uint32                 4                   0 ~ 4 294 967 295
// int64                8             -9 223 372 036 854 775 808 ~ 9 223 372 036 854 775 807
// uint64               8             0 ~ 18 446 744 073 709 551 615
// int              平台相关                    平台相关
// uint             平台相关                    平台相关
// uintptr 同指针 在32位平台下为4字节，64位平台下为8字节



// 需要注意的是，int和int32在Go语言里被认为是两种不同的类型，编译器也不会帮你自动
// 做类型转换，比如以下的例子会有编译错误：
// var value2 int32
// value1 := 64 // value1将会被自动推导为int类型
// value2 = value1 // 编译错误
// 编译错误类似于：
// cannot use value1 (type int) as type int32 in assignment。
// 使用强制类型转换可以解决这个编译错误：
// value2 = int32(value1) // 编译通过
// 当然，开发者在做强制类型转换时，需要注意数据长度被截短而发生的数据精度损失（比如
// 将浮点数强制转为整数）和值溢出（值超过转换的目标类型的值范围时）问题。
