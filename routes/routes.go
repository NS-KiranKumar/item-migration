package routes

import (
	"packages/handlers"
	"packages/middleware"

	"github.com/go-chi/chi/v5"
)

func Routes(userHandler *handlers.UserHandler,
	itemHandler *handlers.ItemHandler,
	rboHandler *handlers.RboHandler) *chi.Mux {

	r := chi.NewRouter()

	r.Post("/Login", userHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Post("/logout", userHandler.Logout)
		r.Post("/ResetPassword/{id}", userHandler.ResetPassword)
		r.Post("/CreateUser", userHandler.CreateUser)
		r.Get("/GetUserbyID/{id}", userHandler.GetUserByID)
		r.Get("/GetAllUsers", userHandler.GetAllUsers)
		r.Delete("/DeleteUserByID/{id}", userHandler.DeleteUserByID)
		r.Post("/CreateItem", itemHandler.CreateItem)
		r.Post("/createRbo", rboHandler.CreateRbo)
		r.Get("/GetRboByID/{id}", rboHandler.GetRbobyID)
		r.Delete("/DeleteRbobyID/{id}", rboHandler.DeleteRbobyID)
		r.Put("/UpdateRbo/{id}", rboHandler.UpdateRbobyID)
		r.Get("/getAllRbos", rboHandler.GetAllRbos)
	})
	return r
}
