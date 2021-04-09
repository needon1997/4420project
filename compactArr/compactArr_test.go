package compactArr_test

import (
	"4420project/compactArr"
	"fmt"
	"github.com/DmitriyVTitov/size"
	"math"
	"testing"
)

var j = 1000
var arr = compactArr.NewCompactArr(32, 1000)
var arr2 = make([]int, 1000)
var arr3 = make(map[int]int)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr3[j] = j
		_ = arr3[j]
	}
}
func TestSize(t *testing.T) {
	fmt.Println(size.Of(make([]int, 1000)))
	fmt.Println(size.Of(compactArr.NewCompactArr(int(math.Log2(float64(1000+1))), 1000)))
}
