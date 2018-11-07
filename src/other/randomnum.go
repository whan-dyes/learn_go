package other

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNum() {
	fmt.Println(">>>>>>>>>>>>>>>> other/RandomNum <<<<<<<<<<<<<<<<<<")
	const N = 10
	var a [N]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(2 * N)

	}

	fmt.Println("rand=", a)

}
