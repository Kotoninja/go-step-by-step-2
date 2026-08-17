package main

import (
	"fmt"
	"sync"
)

type Сount interface {
	Increment()    // увеличение счётчика на единицу
	GetValue() int // получение текущего значения
}

type Counter struct {
	value int
	mu    sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.RLock()
	value := c.value
	c.mu.RUnlock()

	return value
}
