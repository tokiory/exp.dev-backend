package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tokiory/exp.dev-backend/db/report"
	"github.com/tokiory/exp.dev-backend/internal/config"
	"github.com/tokiory/exp.dev-backend/internal/handler"
	"github.com/tokiory/exp.dev-backend/internal/logger"
	"github.com/tokiory/exp.dev-backend/internal/server"
)

func main() {
	log := logger.NewLogger()

	s := server.NewServer(log, server.ServerOptions{
		Addr: ":8080",
	})

	conf := config.NewConfig(log, "config/config.yaml")

	ctx := context.Background()

	pgConf, err := pgxpool.ParseConfig("")
	pgConf.ConnConfig.Host = conf.Database.Host
	pgConf.ConnConfig.Port = conf.Database.Port
	pgConf.ConnConfig.User = conf.Database.User
	pgConf.ConnConfig.Password = conf.Database.Password
	pgConf.ConnConfig.Database = conf.Database.Name

	if err != nil {
		panic("Error parsing database config: " + err.Error())
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgConf)

	if err != nil {
		log.Error("Error connecting to database", slog.String("err", err.Error()))
		return
	}

	reportQuery := report.New(pool)

	handlerOptions := handler.HandlerOptions{
		Log:         log,
		ReportQuery: reportQuery,
	}

	s.Handle("/health", handler.HealthHandler(handlerOptions))
	s.Handle("/report", handler.ReportCreateHandler(handlerOptions))

	s.Apply("Logger", func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("Server: Received a request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
				slog.String("sender", r.RemoteAddr),
			)

			h.ServeHTTP(w, r)
		})
	})

	if err := s.Start(); err != nil {
		log.Error("Error when starting http server", slog.String("err", err.Error()))
	}
}
