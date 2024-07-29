package util

import (
	"math"
	"math/rand"
)

func GenerateRandNum(length int) int {
	min := int(math.Pow(10, float64(length-1)))
	max := int(math.Pow(10, float64(length))) - 1
	return rand.Intn(max-min+1) + min
}
