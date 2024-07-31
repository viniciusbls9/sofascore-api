package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/cors"
)

func HandlerRouters() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", HandlerHealthz)
	v1Router.Post("/create_user", HandlerCreateUser)
	v1Router.Get("/get_user/{userID}", HandlerGetUserByID)
	v1Router.Get("/get_users", HandlerGetUsers)

	router.Mount("/v1", v1Router)

	fmt.Printf("Starting server at port 8000")
	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
