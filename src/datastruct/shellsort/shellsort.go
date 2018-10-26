package shellsort
import "math/rand"
import "time"
import "fmt"
func Run(){
  var x [10]int=[10]int{0:9}
  y:=a(10,x)
  shellsort(y)
}
func shellsort(a []int) {
  fmt.Println("-----------------Shellsort-------------------------------------")
}

func a(n int,b []int) []int {
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
