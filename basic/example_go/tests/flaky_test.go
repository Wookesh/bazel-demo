package tests

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func TestFlaky(t *testing.T) {
	x := rand.Intn(3)
	if x > 1 {
		t.Fail()
	}
}
