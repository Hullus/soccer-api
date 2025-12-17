package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soccer-api/internal/domain/requests"
	"soccer-api/internal/domain/responses"
	"soccer-api/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PlayerHandler struct {
	Service service.TeamService
}

func (h *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	playerID, _ := strconv.ParseInt(chi.URLParam(r, "playerId"), 10, 64)

	var req requests.UpdatePlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdatePlayer(r.Context(), playerID, req.FirstName, req.LastName, req.Country); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(responses.MessageResponse{
		Message: fmt.Sprintf("Player %d updated successfully", playerID),
	})
}
