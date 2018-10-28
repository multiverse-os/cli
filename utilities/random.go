package utilities

import (
	"math/rand"
	"time"
)

func RandomInt(from, to int) int {
	s := rand.NewSource(time.Now().Unix())
	random := rand.New(s)
	return (random.Intn(to) + from)
}
