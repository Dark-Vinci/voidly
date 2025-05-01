package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/handlers"
	"github.com/dark-vinci/stripchat/beetle/utils"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	_ = os.Setenv("TZ", utils.TimeZone)
	ctx := context.Background()

	e := utils.NewEnv()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	appLogger := logger.With().Str("beetle", "api").Logger()

	h := handlers.New(e, logger)
	h.Build(ctx)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", e.Port),
		Handler: h.GetEngine(),
	}

	go func() {
		appLogger.Info().Msgf("server started on %s", e.Port)
		err := server.ListenAndServe()

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			appLogger.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		appLogger.Err(err).Msg("Server forced to shutdown")
	}

	appLogger.Debug().Msg("server last message")
}
