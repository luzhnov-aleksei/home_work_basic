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

func GetUsers(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.UsersHandler.GetUsers"
	data, err := s.GetUsers()
	if err != nil {
		log.Error("Failed to get users", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to get users"))
	}
	render.JSON(w, r, data)
}

func InsertUser(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.UsersHandler.InsertUser"
	var user storage.User

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	json.Unmarshal(b, &user)

	err = s.InsertUser(user)
	if err != nil {
		log.Error("Failed to insert user", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to insert user"))
		return
	}
	responseOK(w, r, "The user has been successfully added")
}

func UpdateUser(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.UsersHandler.UpdateUser"
	var user storage.User

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	json.Unmarshal(b, &user)

	err = s.UpdateUser(user)
	if err != nil {
		log.Error("Failed to update user", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to update user"))
		return
	}
	responseOK(w, r, "The user has been successfully updated")
}

func DeleteUser(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.UsersHandler.DeleteUser"

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	err = s.DeleteUser(string(b))
	if err != nil {
		log.Error("Failed to deleted user", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to deleted user"))
		return
	}
	responseOK(w, r, "The user has been successfully deleted")
}

func GetUserStat(log *slog.Logger, s *storage.Storage, w http.ResponseWriter,
	r *http.Request,
) {
	const op = "handlers.UsersHandler.GetUserStat"
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to read request"))
		return
	}
	defer r.Body.Close()

	userID, _ := strconv.Atoi(string(b))

	sum, avg, err := s.GetUserStat(userID)
	if err != nil {
		log.Error("Failed to get users", fmt.Errorf("%s: %w", op, err))
		render.JSON(w, r, response.Error("Failed to get users"))
	}

	var result struct {
		Sum string `json:"sum"`
		Avg string `json:"avg"`
	}
	result.Sum = fmt.Sprintf("%.2f", sum)
	result.Avg = fmt.Sprintf("%.2f", avg)
	render.JSON(w, r, result)
}
