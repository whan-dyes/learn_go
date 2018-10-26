package generaterandnum
import "math/rand"
import "time"


func GenerateRandnum(n int) int{
  rand.Seed(time.Now().UnixNano())
  rnd := rand.Intn(n)

  return rnd
}
