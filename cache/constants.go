package cache

import (
	"sync"
	"time"
)

// Token to hold bearer token to prevent recreation all the time
type Token struct {
	value       string
	generatedAt time.Time
	generating  bool
	mutex       sync.Mutex
}

// NewToken returns a new token
func NewToken() *Token {
	return &Token{}
}

// GetVal returns the token value
func (c *Token) GetVal() string {
	for {
		c.mutex.Lock()
		if !c.generating {
			c.mutex.Unlock()
			break
		}
		c.mutex.Unlock()
		time.Sleep(1 * time.Second)
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

// SetGenerating sets the token generating status
func (c *Token) SetGenerating() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.generating = true
}

// SetVal sets the token value
func (c *Token) SetVal(val string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = val
	c.generatedAt = time.Now()
	c.generating = false
}
