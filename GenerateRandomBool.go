package gotools

import (
	"math/rand"
	"time"
)

// GenerateRandomBool generates a random boolean value.
func GenerateRandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()%2 == 0
}
