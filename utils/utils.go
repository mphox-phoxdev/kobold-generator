package utils

import "math/rand"

// D6 returns an int between 1 and 6
func D6() int {
	return rand.Intn(6) + 1
}
