package utils

import (
	"math"
)

const base62Symbols = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func EncodeBase62(number int) string {
	base62 := ""
	for number > 0 {
		remainder := number % 62
		base62 = string(base62Symbols[remainder]) + base62
		number /= 62
	}
	return base62
}

func DecodeBase62(base62 string) int {
	id := 0
	power := 0
	for i := len(base62) - 1; i >= 0; i-- {
		pos := getPos(rune(base62[i]))
		id += pos * int(math.Pow(float64(62), float64(power)))
		power++
	}
	return id
}

func getPos(sym rune) int {
	position := 0
	runes := []rune(base62Symbols)
	for i := range len(runes) {
		if sym == runes[i] {
			position = i
			return position
		}
	}
	return -1
}
