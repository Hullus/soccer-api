package handlers

import (
	"encoding/json"
	"net/http"
	"soccer-api/internal/domain/responses"
	"soccer-api/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type MarketHandler struct {
	Service service.MarketService
}

func (h *MarketHandler) ListPlayer(w http.ResponseWriter, r *http.Request) {
	playerID, _ := strconv.ParseInt(chi.URLParam(r, "playerId"), 10, 64)
	var req struct {
		AskingPrice int64 `json:"asking_price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.Service.ListPlayerOnMarket(r.Context(), playerID, req.AskingPrice); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(responses.MessageResponse{
		Message: "Player successfully listed on the transfer market",
	})
}

func (h *MarketHandler) GetMarket(w http.ResponseWriter, r *http.Request) {

	listings, err := h.Service.MarketRepo.GetMarketListings(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch market", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listings)
}

func (h *MarketHandler) CancelListing(w http.ResponseWriter, r *http.Request) {
	playerID, _ := strconv.ParseInt(chi.URLParam(r, "playerId"), 10, 64)
	if err := h.Service.CancelPlayerListing(r.Context(), playerID); err != nil {
		http.Error(w, "failed to Cancel listing", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(responses.MessageResponse{
		Message: "Transfer listing cancelled successfully",
	})
}

func (h *MarketHandler) BuyPlayer(w http.ResponseWriter, r *http.Request) {
	listingID, err := strconv.ParseInt(chi.URLParam(r, "listingId"), 10, 64)
	if err != nil {
		http.Error(w, "invalid listing id", http.StatusBadRequest)
		return
	}

	if err := h.Service.BuyPlayer(r.Context(), listingID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.MessageResponse{
		Message: "Player purchased successfully",
	})
}
