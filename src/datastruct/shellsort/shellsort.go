package shellsort
import randnum  "datastruct/generateRandnum"
import "fmt"
import "time"

const N=10
func Run(){
  var x [N]int

s1:=time.Now()
  y :=  a(N,x)
  s2:=time.Now()
  fmt.Println(s2.Sub(s1))

  t1:=time.Now()
  r:=shellsort(y)
  fmt.Println(r)
  t2:=time.Now()
  fmt.Println(t2.Sub(t1))
}
func shellsort(y [N]int) [N]int{
  fmt.Println("-----------------Shellsort-------------------------------------")

  // fmt.Println(y)
  var i,j,tmp int
  var increment=N

  for increment>1{
    increment=increment/3+1
    for i=increment;i<N;i++{
        if y[i]<y[i-increment]{
        tmp=y[i]
        for j=i-increment;j>=0&&tmp<y[j];j-=increment{
          y[j+increment]=y[j]
        }
        y[j+increment]=tmp
      }
    }
  }
  return y
}

func a(n int,b [N]int) [N]int {
  for  i:=0;i<n;i++ {
    for j:=0;j<1000;j++{
    fmt.Print()
    }
    var x=randnum.GenerateRandnum(10*N)
    b[i]=x
  }
  return b
}
