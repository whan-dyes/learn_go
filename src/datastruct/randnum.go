package datastruct

import (
	// . "fmt"
	"math/rand"
	"time"
)

func randarray(n int) []int {
	b := make([]int, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		b[i] = rand.Intn(5 * N)
	}

	return b
}

func MapRandArray(n int) []int {
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = i + 1
	}

	s := make([]int, len(m))
	i := 0
	for _, v := range m {
		s[i] = v
		i++
	}

	return s
}
