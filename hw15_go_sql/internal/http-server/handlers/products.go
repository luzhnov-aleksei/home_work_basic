package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/luzhnov-aleksei/home_work_basic/hw15_go_sql/internal/lib/api/response"
	"github.com/luzhnov-aleksei/home_work_basic/hw15_go_sql/internal/storage"
	"golang.org/x/exp/slog"
)

func GetProducts(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.ProductsHandler.GetProducts"
	data, err := s.GetProducts()
	if err != nil {
		log.Error("Failed to get products", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to get products"))
	}
	render.JSON(w, r, data)
}

func InsertProduct(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.ProductsHandler.InsertProduct"
	var product storage.Product

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	json.Unmarshal(b, &product)

	err = s.InsertProduct(product)
	if err != nil {
		log.Error("Failed to insert product", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to insert product"))
		return
	}
	responseOK(w, r, "The product has been successfully added")
}

func UpdateProduct(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.ProductsHandler.UpdateProduct"
	var product storage.Product

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	json.Unmarshal(b, &product)

	err = s.UpdateProduct(product)
	if err != nil {
		log.Error("Failed to update product", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to update product"))
		return
	}
	responseOK(w, r, "The product has been successfully updated")
}

func DeleteProduct(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.ProductsHandler.DeleteProduct"

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()
	productID, err := strconv.Atoi(string(b))
	if err != nil {
		log.Error("Failed to decode request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to decode request"))
		return
	}
	err = s.DeleteProduct(productID)
	if err != nil {
		log.Error("Failed to deleted product", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to deleted product"))
		return
	}
	responseOK(w, r, "The product has been successfully deleted")
}
