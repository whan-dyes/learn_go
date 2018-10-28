package shellsort
import "math/rand"
import "time"
import "fmt"
func Run(){

  y:=a()
  shellsort(y)
}
func shellsort(a []int) {
  fmt.Println("-----------------Shellsort-------------------------------------")
}

func a() []int {
  var b [10]int
  n:=10
  for  i:=1;i<=n;i++ {
    b[i]=generateRandnum()
  }
  return b
}

func generateRandnum() int{
  rand.Seed(time.Now().Unix())
  rnd := rand.Intn(100)

  fmt.Printf("rand is %v\n", rnd)

  return rnd
}
