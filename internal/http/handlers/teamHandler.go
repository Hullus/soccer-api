package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soccer-api/internal/service"
)

type TeamHandler struct {
	service service.TeamService
}

func (h *TeamHandler) getTeamInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.service.GetTeamInformation(ctx)
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
