package util

import "math/rand"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func RandomIntBetween(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

func RandomFloatBetween(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return rand.Float64()*(max-min) + min
}
