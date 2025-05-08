package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/thongsoi/architectures/usecase"
)

type RegisterUserRequest struct {
	Name string `json:"name"`
}

type RegisterUserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func RegisterUserHandler(interactor usecase.RegisterUserInteractor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		output, err := interactor.Execute(usecase.RegisterUserInput{Name: req.Name})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(RegisterUserResponse{
			ID:   output.ID,
			Name: output.Name,
		})
	}
}
