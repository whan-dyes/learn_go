package datastruct
import ("math/rand"
 "time"
 . "fmt"
)



func GenerateRandnum(n int) int{
  rand.Seed(time.Now().UnixNano())
  rnd := rand.Intn(n)

  return rnd
}
func randarray(n int) [N]int {
  b:=[N]int{}
  for  i:=0;i<n;i++ {
    for j:=0;j<1000;j++{
    Print()
    }
    var x=GenerateRandnum(5*N)
    b[i]=x
  }
  return b
}

func MapRandArray(n int) []int {
m:=make(map[int]int)
for i:=0;i<n;i++{
  m[i]=i+1
}
// Println(m)
s:=make([]int,len(m))
i:=0
for _,v:=range m {
  s[i]=v
  i++
}
// Println(s)
return s
}
