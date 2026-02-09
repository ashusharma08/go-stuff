package utils

import "math/rand"

func GetRandomNumber(size int) int {
	v := rand.Intn(size)
	return v + 1
}
