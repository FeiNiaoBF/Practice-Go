package sort

import (
	"math/rand"
	"time"
)

var randNums *rand.Rand


func init() {
	seed := time.Now().UnixNano()
	randNums = rand.New(rand.NewSource(seed))
}


func RandNums(num int) (result []int) {
	result = make([]int, num)
	for i := 0; i < num; i++ {
		result[i] = randNums.Intn(1000)
	}
	return
}
