package handlers

import (
	"encoding/json"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/services"
	"net/http"
)

type OrderHandler struct {
	storage *database.Storage
}

func NewOrderHandler(st *database.Storage) *OrderHandler {
	return &OrderHandler{
		storage: st,
	}
}

func (h *OrderHandler) HandleGetAllUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.OrderUID)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		userOrders, err := h.storage.Order().GetAllUserOrders(req.User_id)
		if err != nil {
			services.Error(w, r, http.StatusConflict, err)
			return
		}

		services.Respond(w, r, 200, userOrders)
	}
}
