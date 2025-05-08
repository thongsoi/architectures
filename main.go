package main

import (
	"fmt"
	"net/http"

	"github.com/thongsoi/architectures/repository"
	"github.com/thongsoi/architectures/usecase"

	"github.com/thongsoi/architectures/adapter"
)

func main() {
	userRepo := repository.InMemUserRepository{}
	registerUserUC := usecase.RegisterUserInteractor{UserRepo: userRepo}

	mux := http.NewServeMux()
	mux.HandleFunc("/users", adapter.RegisterUserHandler(registerUserUC))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
