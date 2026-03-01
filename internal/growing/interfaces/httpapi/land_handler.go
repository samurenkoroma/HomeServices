package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"samurenkoroma/services/internal/growing/application/command"
	"samurenkoroma/services/internal/growing/application/query"

	"github.com/google/uuid"
)

type LandUnitHTTPHandler struct {
	createCmd *command.CreateFacilityHandler
	getQuery  *query.GetLandUnitHandler
}

func (h *LandUnitHTTPHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /land-units", h.createLandUnit)
	mux.HandleFunc("GET /land-units/{id}", h.getLandUnit)
}
func NewLandUnitHTTPHandler(
	createCmd *command.CreateFacilityHandler,
	getQuery *query.GetLandUnitHandler,
) *LandUnitHTTPHandler {

	return &LandUnitHTTPHandler{
		createCmd: createCmd,
		getQuery:  getQuery,
	}
}

func (h *LandUnitHTTPHandler) getLandUnit(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	result, err := h.getQuery.Handle(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (h *LandUnitHTTPHandler) createLandUnit(w http.ResponseWriter, r *http.Request) {

	var req CreateFacilityRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := command.CreateFacilityCmd{
		ID:           uuid.New().String(),
		Name:         req.Name,
		Unit:         req.Unit,
		FacilityType: req.Facility,
		Length:       req.Length,
		Width:        req.Width,
	}

	if err := h.createCmd.Handle(context.Background(), cmd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
