package variable
import(
  . "fmt"
)
func Pointer(){
  Println("-------------------pointer------------------")
  test1()
}

func test1(){
  a:=1
  var p *int=&a
  Println("p=",p,"     *p=",*p)
}
func test2(){

}
func test3(){

}
func test4(){

}
func test5(){

}
func test6(){

}
func test7(){

}
/*
指针默认值为 nil 而非 NULL
操作符 & 取变量地址，使用 * 通过指针间接访问目标对象
*/
