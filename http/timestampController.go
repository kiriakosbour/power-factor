package http

import (
	"encoding/json"
	"net/http"
	"power-factor/app"
	"power-factor/domain"
)

type TimestampController struct {
	timestampDataService app.TimestampInteractor
}

func NewTimestampDataHandler(t app.TimestampInteractor) *TimestampController {
	return &TimestampController{
		timestampDataService: t,
	}
}

func (c TimestampController) TimestampsMatching(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	period := r.URL.Query().Get("period")
	tz := r.URL.Query().Get("tz")
	t1 := r.URL.Query().Get("t1")
	t2 := r.URL.Query().Get("t2")

	if period == "" || tz == "" || t1 == "" || t2 == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResp{
			Status: "error",
			Desc:   "invalid/missing parameters",
		})
		return
	}
	timestamp := domain.Timestamp{}
	timestamp = timestamp.GenerateTimestampDomain(period, tz, t1, t2)
	result, err := c.timestampDataService.TimestampService(timestamp)
	if err.Desc != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
