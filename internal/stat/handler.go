package stat

import (
	"net/http"
	"samurenkoroma/services/pkg/response"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.HandleFunc("GET /stat", handler.ShowStat())
}

func (s *StatHandler) ShowStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fromDate, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			response.ErrJson(w, "Invalid date", http.StatusBadRequest)
			return
		}

		toDate, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			response.ErrJson(w, "Invalid date", http.StatusBadRequest)
			return
		}

		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			response.ErrJson(w, "Invalid group param", http.StatusBadRequest)
			return
		}
		response.Json(w, s.StatRepository.GetStats(by, fromDate, toDate), http.StatusOK)
	}
}
