package main

import (
	"fmt"
	"net/http"

	"packages/config"
	"packages/handlers"
	"packages/repository"
	"packages/routes"
)

func main() {

	// 1. Connect MongoDB
	config.ConnectDB()

	// 2. Get users collection
	userCollection := config.DB.Collection("users")
	itemCollection := config.DB.Collection("items")
	rboCollection := config.DB.Collection("rbos")

	// 3. Create Repository
	userRepo := repository.NewUserRepository(userCollection)
	itemRepo := repository.NewItemRepository(itemCollection)
	rboRepo := repository.NewRboRepository(rboCollection)

	// 4. Create Handler
	userHandler := handlers.NewUserHandler(userRepo)
	itemHandler := handlers.NewItemHandler(itemRepo)
	rboHandler := handlers.NewRboHandler(rboRepo)

	// 5. Register Routes
	r := routes.Routes(userHandler, itemHandler, rboHandler)

	// 6. Start Server
	fmt.Println("🚀 Server running on :8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
