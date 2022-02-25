package main

import (
	"log"
	"net/http"

	"github.com/anuragprafulla/bullet/internal/handlers"
	"github.com/anuragprafulla/bullet/internal/users"
	"github.com/gorilla/mux"
)

type Args struct {
	conn string
	port string
}

func Run(args Args) error {
	router := mux.NewRouter().PathPrefix("/api/").Subrouter()
	log.Println(args.conn)
	st := users.NewPostgresUserStore(args.conn)
	hnd := handlers.NewUserHandler(st)

	RegisterAllRoutes(router, hnd)

	log.Println("Starting server at port: ", args.port)

	return http.ListenAndServe(args.port, router)
}

func RegisterAllRoutes(router *mux.Router, hnd handlers.IUserHandler) {

	log.Println("Registered Routes")
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/users", hnd.Get).Methods(http.MethodGet)
	router.HandleFunc("/users", hnd.Create).Methods(http.MethodPost)
	router.HandleFunc("/users", hnd.Delete).Methods(http.MethodDelete)
}
