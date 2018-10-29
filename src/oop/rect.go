package oop
import(
  "fmt"
)

type Rect struct {
  x, y float64
  width, height float64
}

func NewRect(x, y, width, height float64) *Rect {
  return &Rect{x, y, width, height}
}

func (r *Rect) Area() float64 {
  return r.width * r.height
}

func OopTest(){
  rect1 := new(Rect)
  rect2 := &Rect{}
  rect3 := &Rect{0, 0, 100, 200}
  rect4 := &Rect{width: 100, height: 200}
rect5:=NewRect(0,0,10,20)

  fmt.Println("rect1=",rect1)
  fmt.Println("rect2=",rect2)
  fmt.Println("rect3=",rect3)
  fmt.Println("rect4=",rect4)
  fmt.Println("rect5=",rect5)
  
  fmt.Println(rect5.Area())
}
