package domain

import (
	"fmt"
	"math/rand"
	"time"
)

// ID64 implements ID interface
type ID64 int64

var rnd *rand.Rand

// NewID64 generates a random 64 bits
func NewID64() ID64 {
	if rnd == nil {
		rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return ID64(rnd.Int63())
}

func (id ID64) String() string {
	return fmt.Sprintf("%x", int64(id))
}
