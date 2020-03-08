package concurrent

import "sync"

type StringSlice struct {
	s   []string
	len int
	cap int

	mu           sync.RWMutex
	pendingClear bool
}

func NewStringSlice(cap int) *StringSlice {
	return &StringSlice{
		s:   make([]string, cap),
		len: 0,
		cap: cap,
	}
}

func (s *StringSlice) Add(e string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	next := s.Len()
	if next >= cap(s.s) {
		return false
	}

	s.s[next] = e
	s.len += 1
	return true
}

func (s *StringSlice) Cap() int {
	return s.cap
}

func (s *StringSlice) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.actuallyClear()
}

func (s *StringSlice) actuallyClear() {
	s.s = make([]string, s.cap)
	s.len = 0
}

func (s *StringSlice) Len() int {
	return s.len
}

// Read makes the slice available for reading and only reading. Do not modify the slice.
// Set clear to true to clear after ReadDone is called
// ReadDone must be called when finished to release the read lock.
func (s *StringSlice) Read(clear bool) []string {
	if clear {
		s.mu.Lock()
		s.pendingClear = true
	} else {
		s.mu.RLock()
	}
	return s.s[:s.len]
}

// ReadDone releases a read lock after Read is called.
func (s *StringSlice) ReadDone() {
	if s.pendingClear {
		s.actuallyClear()
		s.pendingClear = false
		s.mu.Unlock()
	} else {
		s.mu.RUnlock()
	}
}
