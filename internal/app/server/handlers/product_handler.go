package handlers

import (
	"encoding/json"
	"github.com/Heroin-lab/nixProject/models"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/services"
	"net/http"
)

type ProductHandler struct {
	resService *services.RespondService
	storage    *database.Storage
}

func NewProductHandler(st *database.Storage) *ProductHandler {
	return &ProductHandler{
		storage: st,
	}
}

func (h *ProductHandler) HandleGetProductsByCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var req = new(models.CategoryRequest)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		getItems, err := h.storage.Product().GetByCategory(req.Category_name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		h.resService.Respond(w, r, 200, getItems)
	}
}

func (h *ProductHandler) HandleInsertProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var req = new(models.Products)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		insertItem, err := h.storage.Product().InsertItem(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		h.resService.Respond(w, r, 200, insertItem)
	}
}

func (h *ProductHandler) HandleDeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var req = new(models.Products)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := h.storage.Product().DeleteItem(req.Product_name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		h.resService.Respond(w, r, 200, "Delete was successfully made")
	}
}

func (h *ProductHandler) HandleUpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var req = new(models.Products)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := h.storage.Product().UpdateItem(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		h.resService.Respond(w, r, 200, "Update was successfully made")
	}
}
