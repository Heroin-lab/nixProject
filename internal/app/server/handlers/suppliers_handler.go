package handlers

import (
	"encoding/json"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/services"
	"net/http"
)

type SuppliersHandler struct {
	resService *services.RespondService
	storage    *database.Storage
}

func NewSuppliersHandler(st *database.Storage) *SuppliersHandler {
	return &SuppliersHandler{
		storage: st,
	}
}

func (h *SuppliersHandler) HandleGetSuppliersByCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		req := new(models.Suppliers)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			h.resService.Error(w, r, http.StatusBadRequest, err)
			return
		}

		result, err := h.storage.Supplier().GetSuppliersByCategory(req.Type)
		if err != nil {
			h.resService.Error(w, r, http.StatusNotFound, err)
		}

		h.resService.Respond(w, r, 200, result)
	}
}
