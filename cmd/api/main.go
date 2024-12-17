package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ExonegeS/AP1_Assignment1/internal/library"
	"github.com/ExonegeS/AP1_Assignment1/internal/shapes"
)

func main() {
	ctx := context.Background()

	config, err := getConfig(ctx)
	if err != nil {
		log.Fatalf("error getConfig: %v", err)
	}

	var (
		logger = getLogger("terminal")
		mux    = http.NewServeMux()
	)

	librarySvc := library.NewService()

	shapesSvc := shapes.NewService()

	mux.Handle("/library/", library.NewHandler(librarySvc, logger))
	mux.Handle("/shapes/", shapes.NewHandler(shapesSvc, logger))

	http.Handle("/", mux)
	run(config, *logger)
}

func run(config *config, logger slog.Logger) {
	addr := ":" + config.Port

	errs := make(chan error, 2)
	go func() {
		logger.Info("listening", "transport", "http", "address", addr)
		errs <- http.ListenAndServe(addr, nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Error("terminated", "errs", <-errs)
}

func getLogger(outputType string) *slog.Logger {
	var handler slog.Handler

	switch outputType {
	case "file":
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			slog.Error("Failed to open log file: %v", "err", err)
			os.Exit(1)
		}
		handler = slog.NewTextHandler(file, &slog.HandlerOptions{Level: slog.LevelDebug})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	logger := slog.New(handler)

	return logger
}
