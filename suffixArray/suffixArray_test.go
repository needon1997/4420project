package suffixArray_test

import (
	"4420project/gvcsa"
	"4420project/suffixArray"
	"4420project/util"
	"math"
	"testing"
)

var alphabet = 4
var logn float64 = 20
var str = util.GenRandomStr(int(math.Pow(2, logn)), alphabet)

var sa2 = suffixArray.CreateSuffixArray(str)
var wtfmi = sa2.ToWTFMI()
var rlfmi = sa2.ToRLFMI()
var gv = gvcsa.MakeGVCSArray(*sa2)
var sad = gvcsa.MakeSADCSArray(*sa2)
var l = len(str)

func BenchmarkNewSuffixArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		suffixArray.CreateSuffixArray(str)
	}
}

func BenchmarkNewGVCSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gvcsa.MakeGVCSArray(*suffixArray.CreateSuffixArray(str))
	}
}
func BenchmarkNewSADCSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gvcsa.MakeSADCSArray(*suffixArray.CreateSuffixArray(str))
	}
}
func BenchmarkNewWTFMI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		suffixArray.CreateSuffixArray(str).ToWTFMI()
	}
}

func BenchmarkNewRLFMI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		suffixArray.CreateSuffixArray(str).ToRLFMI()
	}
}

func BenchmarkWTFMI_Locate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % l
		wtfmi.Locate(j)
	}
}
func BenchmarkRLFMI_Locate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % l
		rlfmi.Locate(j)
	}
}
func BenchmarkSuffixArray_Locate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % l
		_ = sa2.POS[j]
	}
}
func BenchmarkGV_Locate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % l
		gv.Lookup(j)
	}
}
func BenchmarkSAD_Locate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % l
		sad.Lookup(j)
	}
}

func BenchmarkWTFMI_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			str1 := util.GenRandomStr(20, alphabet)
			wtfmi.Search(str1)
		} else {
			j := i % (l - 20)
			wtfmi.Search(str[j : j+20])
		}
	}
}
func BenchmarkRLFMI_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			str1 := util.GenRandomStr(20, alphabet)
			rlfmi.Search(str1)
		} else {
			j := i % (l - 20)
			rlfmi.Search(str[j : j+20])
		}
	}
}
func BenchmarkSuffixArray_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			str1 := util.GenRandomStr(20, alphabet)
			sa2.Search(str1)
		} else {
			j := i % (l - 20)
			sa2.Search(str[j : j+20])
		}
	}
}
func BenchmarkGV_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			str1 := util.GenRandomStr(20, alphabet)
			gv.Search(str1)
		} else {
			j := i % (l - 20)
			gv.Search(str[j : j+20])
		}
	}
}
func BenchmarkSAD_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			str1 := util.GenRandomStr(20, alphabet)
			sad.Search(str1)
		} else {
			j := i % (l - 20)
			sad.Search(str[j : j+20])
		}
	}
}

func BenchmarkWTFMI_Substring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % (l - 20)
		wtfmi.Substring(j, j+20)
	}
}
func BenchmarkRLFMI_Substring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % (l - 20)
		rlfmi.Substring(j, j+20)
	}
}
func BenchmarkWTFMI_Substring2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := i % (l - 20)
		sad.Substring(j, j+20)
	}
}
