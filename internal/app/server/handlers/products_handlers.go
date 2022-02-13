package handlers

import (
	"encoding/json"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/services"
	"net/http"
)

type ProductHandler struct {
	storage *database.Storage
}

func NewProductHandler(st *database.Storage) *ProductHandler {
	return &ProductHandler{
		storage: st,
	}
}

func (h *ProductHandler) HandleGetProductsByCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.CategoryRequest)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		getItems, err := h.storage.Product().GetByCategory(req.Category_name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		services.Respond(w, r, 200, getItems)
	}
}

func (h *ProductHandler) HandleInsertProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = new(models.Products)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		insertItem, err := h.storage.Product().InsertItem(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		services.Respond(w, r, 200, insertItem)
	}
}

func (h *ProductHandler) HandleDeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = new(models.Products)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		err := h.storage.Product().DeleteItem(req.Product_name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		services.Respond(w, r, 200, "Delete was successfully made")
	}
}

func (h *ProductHandler) HandleUpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = new(models.Products)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		err := h.storage.Product().UpdateItem(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		services.Respond(w, r, 200, "Update was successfully made")
	}
}
