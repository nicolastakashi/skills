// Package registry maps encoder names to encoders.
package registry

import (
	"errors"
	"strconv"
)

type Encoder func(string) string

var ErrDuplicate = errors.New("registry: name already registered")

var encoders = map[string]Encoder{}

// Register adds an encoder. A name can be registered only once.
func Register(name string, e Encoder) error {
	if _, ok := encoders[name]; ok {
		return ErrDuplicate
	}
	encoders[name] = e
	return nil
}

// Lookup returns the encoder registered under name.
func Lookup(name string) (Encoder, bool) {
	e, ok := encoders[name]
	return e, ok
}

// Quote is a ready-made encoder that wraps a value in double quotes.
func Quote(s string) string { return strconv.Quote(s) }
