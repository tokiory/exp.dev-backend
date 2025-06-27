package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tokiory/exp.dev-backend/db/report"
	"github.com/tokiory/exp.dev-backend/internal/model"
)

func ReportCreateHandler(opts HandlerOptions) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		req := new(model.ReportAddReq)

		if err := decoder.Decode(req); err != nil {
			opts.Log.Error("Failed to decode request body", "error", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		uuid, err := opts.ReportQuery.CreateReport(ctx)

		if err != nil {
			opts.Log.Error("Failed to create report", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var wg sync.WaitGroup

		wg.Add(3)
		errCh := make(chan error, 3)

		go func() {
			defer wg.Done()
			var patronymic pgtype.Text

			if err := patronymic.Scan(req.Person.Patronymic); err != nil {
				opts.Log.Error("Failed to parse patronymic", "error", err)
				errCh <- err
				return
			}

			if err := opts.ReportQuery.CreateReportPerson(ctx, report.CreateReportPersonParams{
				ReportID:   uuid,
				Name:       req.Person.Name,
				Surname:    req.Person.Surname,
				Patronymic: patronymic,
				Email:      req.Person.Email,
				Telegram:   req.Person.Telegram,
			}); err != nil {
				opts.Log.Error("Failed to create report person", "error", err)
				errCh <- err
			}
		}()

		go func() {
			defer wg.Done()
			if err := opts.ReportQuery.CreateReportWork(ctx, report.CreateReportWorkParams{
				ReportID: uuid,
				Position: req.Work.Position,
				Grade: req.Work.Grade,
				GrowthMessage: req.Work.GrowthMessage,
				TasksMessage: req.Work.TasksMessage,
			}); err != nil {
				opts.Log.Error("Failed to create report work entity", "error", err)
				errCh <- err
			}
		}()

		go func() {
			defer wg.Done()
			jsonSkills, err := json.Marshal(req.Skills)

			if err != nil {
				opts.Log.Error("Failed to marshal skill map", "error", err)
				errCh <- err
				return
			}

			if err := opts.ReportQuery.CreateReportSkills(ctx, report.CreateReportSkillsParams{
				ReportID: uuid,
				Skills: jsonSkills,
			}); err != nil {
				opts.Log.Error("Failed to create report skills map", "error", err)
				errCh <- err
			}
		}()

		// Wait for all goroutines to complete
		wg.Wait()
		close(errCh)

		// Check if any errors occurred
		for err := range errCh {
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}

		res := model.ReportAddRes{
			Id: uuid.String(),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			opts.Log.Error("Failed to encode response", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})
}
