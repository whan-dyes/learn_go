package function
import(
  . "fmt"
)

func MyDefer(){
    Println("<<<<<<<<<<<<<<<<<<<<< MyDefer >>>>>>>>>>>>>>>>>>>>>>>>")
    defer1()
    defer2()
    defer3()
    defer4()
}

func defer1(){
  Println("............ defer1() ...........")
}

func defer2(){
  Println("............ defer2() ...........")
  defer func(){
    if err:=recover();err !=nil{
      Println("Recover in defer2")
    }
  }()
  panic("Panic in defer2")
}

func defer3(){
  Println("............ defer3() ...........")
}

func defer4(){
  Println("............ defer4() ...........")
  var fs=[4]func(){}

  for i:=0;i<4;i++{
    defer Println("defer i=",i)
    defer func(){Println("defer_closure i=",i)}()
    fs[i]=func(){Println("closure i=",i)}
  }

  for _,f :=range fs{
    f()
  }
}
