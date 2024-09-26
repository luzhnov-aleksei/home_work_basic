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

func GetOrders(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.OrderHandler.GetOrder"

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	userID, _ := strconv.Atoi(string(b))
	data, err := s.GetOrders(userID)
	if err != nil {
		log.Error("Failed to get orders", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to get orders"))
	}
	render.JSON(w, r, data)
}

func InsertOrder(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.OrderHandler.InsertOrder"
	var order storage.NewOrder

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	json.Unmarshal(b, &order)

	err = s.InsertOrder(order)
	if err != nil {
		log.Error("Failed to insert order", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to insert order"))
		return
	}
	responseOK(w, r, "The order has been successfully added")
}

func DeleteOrder(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.OrderHandler.DeleteOrder"

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	orderID, _ := strconv.Atoi(string(b))
	err = s.DeleteOrder(orderID)
	if err != nil {
		log.Error("Failed to deleted order", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to deleted order"))
		return
	}
	responseOK(w, r, "The order has been successfully deleted")
}
