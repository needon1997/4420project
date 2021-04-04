package util

import "math/rand"

func GenRandomStr(length int, alphabet int) string {
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		str[i] = byte(rand.Float64()*float64(alphabet) + 97)
	}
	return string(str)
}

func GenRandomStrRepeat(length int, alphabet int, repeated int) string {
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		str[i] = byte(rand.Float64()*float64(alphabet) + 97)
	}
	result := ""
	for i := 0; i < repeated; i++ {
		result += string(str)
	}
	return result
}
