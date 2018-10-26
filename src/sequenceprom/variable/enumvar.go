package variable
import "fmt"

func EnumVar(){
  //   枚举指一系列相关的常量，比如下面关于一个星期中每天的定义。通过上一节的例子，我们
  // 看到可以用在const后跟一对圆括号的方式定义一组常量，这种定义法在Go语言中通常用于定义
  // 枚举值。Go语言并不支持众多其他语言明确支持的enum关键字。
  // 下面是一个常规的枚举表示法，其中定义了一系列整型常量：
  const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday=iota
    Friday
    Saturday
    numberOfDays // 这个常量没有导出
  )
  fmt.Println("Saturday=",Saturday)
  fmt.Println("Sunday=",Sunday)
  // 同Go语言的其他符号（symbol）一样，以大写字母开头的常量在包外可见。
  // 以上例子中numberOfDays为包内私有，其他符号则可被其他包访问。
}
