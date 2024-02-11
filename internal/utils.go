package internal

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomBool() bool {
	return rand.Intn(100) < 10
}
