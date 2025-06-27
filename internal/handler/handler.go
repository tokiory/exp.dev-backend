package handler

import (
	"log/slog"
	"net/http"

	"github.com/tokiory/exp.dev-backend/db/report"
)

type HandlerOptions struct {
	Log *slog.Logger
	ReportQuery *report.Queries
}

type ServerHandler func (HandlerOptions) http.HandlerFunc
