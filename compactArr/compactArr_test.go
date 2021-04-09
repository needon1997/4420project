package compactArr_test

import (
	"4420project/compactArr"
	"testing"
)

var j = 1000
var arr = compactArr.NewCompactArr(32, 1000)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := i % j
		arr.Get(k)
	}
}
