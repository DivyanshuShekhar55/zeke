package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

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

	var cfg config

	// take in info from terminal at server start time
	flag.StringVar(&cfg.addr, "addr", envOrVar("ZEKE_ADDR", ":6969"), "API server address")
	flag.StringVar(&cfg.env, "env", envOrVar("ZEKE_ENV", "dev"), "environment mode for address")
	flag.Parse()

	// handle SIGINT gracefully (CTRL + C)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// handle shutdown properly, ensure no leaks, close everything here
	defer func() {
		// pass for now
	}()

	// create the app and create the net listener
	app := &app{
		config: cfg,
	}

	srv := &http.Server{
		Addr:         cfg.addr,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Minute * 2,
		IdleTimeout:  time.Minute * 1,
		BaseContext:  func(net.Listener) context.Context { return ctx },
		Handler:      app.routes(),
	}

	// handle interrupts
	srvErr := make(chan error, 1)
	go func(){
		srvErr <- srv.ListenAndServe()
	}()

	return nil
}

func(app *app) routes() http.Handler {
	mux:= http.NewServeMux()

	mux.HandleFunc("/upload", nil)

	return mux
}

func envOrVar(key, defaultVal string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultVal
}
