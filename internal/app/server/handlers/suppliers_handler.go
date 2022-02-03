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
	storage *database.Storage
}

func NewSuppliersHandler(st *database.Storage) *SuppliersHandler {
	return &SuppliersHandler{
		storage: st,
	}
}

func (h *SuppliersHandler) HandleGetSuppliersByCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.Suppliers)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		itemSet, err := h.storage.Supplier().GetSuppliersByCategory(req.Type)
		if err != nil {
			services.Error(w, r, http.StatusNotFound, err)
		}

		services.Respond(w, r, 200, itemSet)
	}
}

func (h *SuppliersHandler) HandleAddSupplier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.Suppliers)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		newItem, err := h.storage.Supplier().AddSupplier(req)
		if err != nil {
			services.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		services.Respond(w, r, http.StatusCreated, newItem)
	}
}

func (h *SuppliersHandler) HandleDeleteSupplier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.Suppliers)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := h.storage.Supplier().DeleteSupplier(req.Id); err != nil {
			services.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		services.Respond(w, r, http.StatusOK, "Supplier with id='"+req.Id+"' was successfully deleted!")
	}
}

func (h *SuppliersHandler) HandleUpdateSupplier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.Suppliers)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		err := h.storage.Supplier().UpdateSupplier(req)
		if err != nil {
			services.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		services.Respond(w, r, http.StatusOK, "Supplier with id='"+req.Id+"' was successfully updated!")
	}
}
