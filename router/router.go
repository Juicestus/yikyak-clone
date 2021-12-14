package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRouter() *mux.Router {

	router := mux.NewRouter()

	// Host the react app build
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./client/build/"))))

	/*
	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	*/

	return router
}