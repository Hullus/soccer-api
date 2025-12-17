package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soccer-api/internal/domain/requests"
	"soccer-api/internal/service"
)

type TeamHandler struct {
	Service service.TeamService
}

func (h *TeamHandler) GetTeamInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.Service.GetTeamInformation(ctx)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h *TeamHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var req requests.UpdateTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateTeam(r.Context(), req.Name, req.Country); err != nil {
		http.Error(w, "failed to update team", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
