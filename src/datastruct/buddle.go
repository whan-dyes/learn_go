package datastruct

import(
  . "fmt"
)

func Buddle(){
  a:=[...]int{1,4,9,2,8,9,3}
  Println("a=",a)

  num:=len(a)
  for i:=0;i<num;i++{
    for j:=i+1;j<num;j++{
      if a[i]>a[j]{
        temp:=a[i]
        a[i]=a[j]
        a[j]=a[i]
      }
    }
  }
  Println("sorted a=",a)
}
