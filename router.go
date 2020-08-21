package main

import (
	"url-shortener/links"
	"url-shortener/middleware"
	"url-shortener/users"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// LoadRouter for handle all router
func LoadRouter(db *gorm.DB) (r *mux.Router) {
	r = mux.NewRouter()
	r.Use(middleware.LoggerMiddleware)

	v1 := r.PathPrefix("/api/v1").Subrouter()

	userRouter(v1, db)
	linkRouter(v1, db)
	shortenRouter(r, db)

	return
}

func userRouter(route *mux.Router, db *gorm.DB) {
	userRepo := users.NewUsersRepo(db)
	userController := users.NewUsersController(userRepo)

	route.HandleFunc("/users", userController.Resources).Methods("GET", "POST")
	route.HandleFunc("/users/{id}", userController.Resources).Methods("GET", "PUT", "DELETE")
	route.HandleFunc("/signin", userController.SignIn).Methods("POST")
	route.HandleFunc("/refresh", userController.RefreshToken).Methods("GET")
}

func linkRouter(route *mux.Router, db *gorm.DB) {
	userRepo := users.NewUsersRepo(db)
	linkRepo := links.NewLinksRepo(db)
	linkController := links.NewLinksController(linkRepo, userRepo)

	route.HandleFunc("/links", linkController.Resources).Methods("POST")
	route.HandleFunc("/links/{id}", linkController.Resources).Methods("GET", "PUT", "DELETE")
}

func shortenRouter(route *mux.Router, db *gorm.DB) {
	userRepo := users.NewUsersRepo(db)
	linkRepo := links.NewLinksRepo(db)
	linkController := links.NewLinksController(linkRepo, userRepo)

	route.HandleFunc("/l/{translated_url:[a-zA-Z0-9]{5}}", linkController.Shorten).Methods("GET")
}
