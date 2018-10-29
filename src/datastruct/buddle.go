package datastruct

import(
  . "fmt"
  "time"
)

func Buddle(){
  Println("-----------------Buddle---------------------------------------------")
  a:=MapRandArray(N)
  t1:=time.Now()
  buddle1(a)
  t2:=time.Now()
  Println("buddle run time: ",t2.Sub(t1))
}

func buddle1(a []int){
    Println("start  a=",a)
  num:=len(a)
  for i:=0;i<num;i++{
    for j:=i+1;j<num;j++{
      if a[i]>a[j]{
        temp:=a[i]
        a[i]=a[j]
        a[j]=temp
      }
    }
  }
  Println("sorted a=",a)
}
