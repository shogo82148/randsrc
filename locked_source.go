package randsrc

import (
	"math/rand"
	"sync"
)

type lockedSource struct {
	mu  sync.Mutex
	src rand.Source
}

func (s *lockedSource) Int63() (n int64) {
	s.mu.Lock()
	n = s.src.Int63()
	s.mu.Unlock()
	return
}

func (s *lockedSource) Seed(seed int64) {
	s.mu.Lock()
	s.src.Seed(seed)
	s.mu.Unlock()
	return
}

type lockedSource64 struct {
	mu  sync.Mutex
	src rand.Source64
}

func (s *lockedSource64) Int63() (n int64) {
	s.mu.Lock()
	n = s.src.Int63()
	s.mu.Unlock()
	return
}

func (s *lockedSource64) Uint64() (n uint64) {
	s.mu.Lock()
	n = s.src.Uint64()
	s.mu.Unlock()
	return
}
func (s *lockedSource64) Seed(seed int64) {
	s.mu.Lock()
	s.src.Seed(seed)
	s.mu.Unlock()
}

// NewLockedSource returns wrapped src that is safe for concurrent use.
func NewLockedSource(src rand.Source) rand.Source {
	if s64, ok := src.(rand.Source64); ok {
		return &lockedSource64{src: s64}
	}
	return &lockedSource{src: src}
}
