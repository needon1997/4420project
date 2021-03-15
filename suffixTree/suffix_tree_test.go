package suffixTree_test

import (
	"fmt"
	"testing"
)
import suffix "text-indexing"

func TestBuild(t *testing.T) {
	tree := suffix.NewSuffixTree("abcabxabcd")
	arr := tree.ToSuffixArray()
	fmt.Println(arr.POS)
	fmt.Println(arr.LCP)
	arr.Search("a")
}
func BenchmarkBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := suffix.NewSuffixTree("yH823YG7f3xVmAXqE40lc1346reXDMfc0WE6cHmnGUkd6wWG2tq9Rgi9mXcQx83DRv6Zi7VOCM6v6ws2sXhDZo0wog3tiAzo9h79ZqJA2V8uvkYK1wp7RvHr1Q6EBNHAa2OKV3fGz8V0Dz3EucA9qwa7unLfPhG6vMuS5M797jpyOhp7fhi8e1DEqGiPnDm7k6M2O61zvlRdQEoBw30zEGK2y2fHXfvg28a6gYXPjvz4Vw76YgQMqw0dz7kdy17i88E6TIFob5JUBn06qWm6038ZJ7s7u70xLMsoE8HmkceH8y277qC2iQRSXXitoikF393668FXZI3Z1qC7sy7n9795QJ39vqT797Y91yLI7yfn92i0pMhQjP01qCQV3Py2FfQ3o06qQm36WzxXFAsxtbd25fF0dp98WO55wu7S1MAfs26qbojZHxt6AGyQedUOVWnF6SHwI9qLdHigkUdU2NQFoA3849gRwsgHTrFlnl3ol842VErx7qetjEa77mTB8843z65vhH94ivmT2Z8thknAIalFt46ma5ZGKNK5tX90Hys725SPS75")
		tree.ToSuffixArray()
	}
}

func TestRadix(t *testing.T) {
	//arr := []string{"abc","cde","aaa","bbb","123","jkl"}
	//barr := suffix.ToByte(arr)
	//suffix.Radixsort(&barr,3)
	t3 := suffix.Build("helloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworld")
	t3_b := suffix.ToByte(t3)
	suffix.Radixsort(&t3_b, 3)
	fmt.Println(t3)
}
