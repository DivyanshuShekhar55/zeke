package main

import "log"

type app struct {
	config config
}

type config struct {
	addr string // net listener address
	env  string // dev or prod
}

func main() {
	if err := run(); err != nil {
		log.Fatal("Error: Could not start server", err)
	}
}

func run() error {
	return nil
}
