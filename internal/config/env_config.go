package config

import (
	"os"
	"strconv"
	"time"
)

var (
	AllowedSkew time.Duration
)

func Load() {
	const defaultSkewSeconds = 10

	val := os.Getenv("GYROSCOPE_ALLOWED_SKEW_SECONDS")
	if val == "" {
		AllowedSkew = time.Duration(defaultSkewSeconds) * time.Second
		return
	}

	seconds, err := strconv.Atoi(val)
	if err != nil || seconds < 0 {
		AllowedSkew = time.Duration(defaultSkewSeconds) * time.Second
		return
	}

	AllowedSkew = time.Duration(seconds) * time.Second
}
