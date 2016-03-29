package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"src/github.com/geoffroya/stdin-beat/stdin-beat/beater"
)

func main() {
	err := beat.Run("stdin-beat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
