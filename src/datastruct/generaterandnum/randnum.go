package generaterandnum
import "math/rand"
import "time"
// var(
//   Name int=300000
// )
// const PI=3.14
func GenerateRandnum(n int) int{
  rand.Seed(time.Now().UnixNano())
  rnd := rand.Intn(n)

  return rnd
}
