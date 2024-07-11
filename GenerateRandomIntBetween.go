package gotools

import (
	"math/rand"
	"time"
)

// GenerateRandomIntBetween generates a random integer between min and max.
func GenerateRandomIntBetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
