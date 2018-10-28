package function

import(
  . "fmt"
)

func MyFunc(){
  Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>> MyFunc <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
  myfunc1()
  myfunc2()
  s:=[]int{1,2,3}
  myfunc3(s,4,5,6,7,8)
  Println(s)

  a:=1
  Println("Myfunc() a=",a)
  M:=myfunc4      //golang中函数也是一种类型
  M(&a)
  Println("Myfunc() a=",a)
}

func myfunc1()(a,b,c int){
  Println("............ myfunc1() ...........")
  a,b,c=1,2,3
  return
}

func myfunc2()(int,int,int){
  Println("............ myfunc2() ...........")
  a,b,c:=1,2,3
  return a,b,c
}

func myfunc3(b []int,a ...int) []int{     // ...int 相当于 slice，但是值传递，不是引用
  Println("............ myfunc3() ...........")
  Println(b)
  b[0]=100

  Println(a)
  a[0]=100

  return a
}

func myfunc4(a *int){
  Println("............ myfunc4() ...........")
  *a=100
  Println("*a=",*a)
}
