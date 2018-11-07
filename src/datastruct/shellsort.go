package datastruct

import "fmt"
import "time"

func Run() {
	// var x [N]int

	s1 := time.Now()
	y := randarray(N)
	s2 := time.Now()
	fmt.Println(s2.Sub(s1))

	t1 := time.Now()
	r := shellsort(y)
	fmt.Println(r)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}

func shellsort(y []int) []int {
	fmt.Println("-----------------Shellsort-------------------------------------")

	var i, j, tmp int
	var increment = len(y)

	for increment > 1 {
		increment = increment/3 + 1
		for i = increment; i < N; i++ {
			if y[i] < y[i-increment] {
				tmp = y[i]
				for j = i - increment; j >= 0 && tmp < y[j]; j -= increment {
					y[j+increment] = y[j]
				}
				y[j+increment] = tmp
			}
		}
	}
	return y
}
