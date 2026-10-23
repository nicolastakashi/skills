// Package config loads server settings.
package config

import (
	"fmt"
	"strconv"
	"time"
)

// Config holds server settings.
type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
}

// Raw is the config as read from the file.
type Raw struct {
	Host string
	Port string
	// TimeoutSeconds is the request timeout in seconds.
	TimeoutSeconds int
}

// FromRaw validates raw and converts it to a Config.
func FromRaw(raw Raw) (Config, error) {
	port, err := ParsePort(raw.Port)
	if err != nil {
		return Config{}, err
	}
	return Config{
		Host:    raw.Host,
		Port:    port,
		Timeout: time.Duration(raw.TimeoutSeconds) * time.Millisecond,
	}, nil
}

// ParsePort parses a TCP port. Valid ports are 1 to 65535.
func ParsePort(s string) (int, error) {
	p, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("port %q: %w", s, err)
	}
	if p < 1 || p > 65536 {
		return 0, fmt.Errorf("port %d out of range", p)
	}
	return p, nil
}
