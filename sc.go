package main

import (
	"fmt"
	"sync"
)

var (
	container   map[string]interface{}
	containerMu sync.RWMutex
)

func init() {
	container = make(map[string]interface{})
}

func getTypeKey[T any](input T) string {
	return fmt.Sprintf("%T", input)
}

// Put puts an value in the container.
func Put[T any](input T) {
	typeKey := getTypeKey(input)

	containerMu.Lock()
	defer containerMu.Unlock()
	container[typeKey] = input
}

// Get gets the value from the container.
func Get[T any]() (T, bool) {
	var input T
	typeKey := getTypeKey(input)

	containerMu.RLock()
	defer containerMu.RUnlock()
	res, found := container[typeKey]

	if !found {
		return input, false
	}

	return res.(T), true
}

