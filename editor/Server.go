package editor

import (
	"context"
	mux "github.com/hyperjumptech/hyper-mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	hmux = mux.NewHyperMux()
)

func Start() {
	InitializeStaticRoute(hmux)
	InitializeEvaluationRoute(hmux)

	var wait time.Duration
	address := "0.0.0.0:32123"

	logrus.Infof("Starting Editor at http://%s", address)

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      hmux, // Pass our instance of gorilla/mux in.
	}

	startTime := time.Now()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	channel := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(channel, os.Interrupt)

	// Block until we receive our signal.
	<-channel

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	dur := time.Now().Sub(startTime)
	logrus.Infof("Shutting down. This Satpam been protecting the world for %s", dur.String())
	os.Exit(0)

}
