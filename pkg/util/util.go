package util

import (
	"math/rand"
	"time"
)

func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
