package gotools

import (
	"math/rand"
	"time"
)

// GenerateRandomInt generates a random integer.
func GenerateRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
