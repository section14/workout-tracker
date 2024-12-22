package api

import (
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /exercises", getExercises)
	mux.HandleFunc("GET /exercise/{id}", getExercises)
	http.ListenAndServe(":8080", mux)
}

func getExercises(w http.ResponseWriter, r *http.Request) {

}

func getExercise(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
    fmt.Println(idStr)
}
