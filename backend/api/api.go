package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jub0bs/cors"
	"github.com/section14/workout-tracker/backend/store"
)

// reference to global datastore
var DataStore store.Store

func LoadStore() error {
	storeName := os.Args[1]

	date := time.Now()
	log.Println(date)

	file, err := os.Open(storeName)
	if err != nil {
		return errors.New(fmt.Sprintf("couldn't open store: %s", storeName))

	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&DataStore)
	if err != nil {
		//decodeErr := errors.New(fmt.Sprintf("couldn't decode store: %s", storeName))
		//return store.Store{}, decodeErr
	}

    log.Println(DataStore)

	//validate and parse supplied JSON file
	/*
		jsonData, err := validateJson(storeName)
		if err != nil {
			log.Fatal(err)
		}
	*/

	return nil
}

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /exercises", getExercises)
	mux.HandleFunc("GET /exercise/{id}", getExercise)
    mux.HandleFunc("POST /exercise", createExercise)
	mux.HandleFunc("DELETE /exercise/{id}", deleteExercise)
	mux.HandleFunc("PATCH /exercise/{id}", updateExercise)
}

func Serve() {
	mux := http.NewServeMux()

	//handlers
    handlers(mux)

	//setup
	corsMiddleware, err := cors.NewMiddleware(cors.Config{
		Origins: []string{"*"},
		Methods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
		},
		RequestHeaders: []string{
			"Authorization",
			"Content-Type",
			"X-Requested-With",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	s := &http.Server{
		Addr:    "localhost:8080",
		Handler: corsMiddleware.Wrap(mux),
		// other settings omitted
	}
	fmt.Println("Serving on port 8080")
	log.Fatal(s.ListenAndServe())
	//http.ListenAndServe(":8080", mux)
}

func getExercises(w http.ResponseWriter, r *http.Request) {
	res, err := DataStore.Exercises.GetAll()
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusNotFound)
    }
	w.Write(res)
}

func getExercise(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)

    res, err := DataStore.Exercises.Get(id)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusNotFound)
    }
    w.Write(res)
}

func createExercise(w http.ResponseWriter, r *http.Request) {
    err := DataStore.Exercises.Create(r)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusNotFound)
    }

    w.WriteHeader(http.StatusOK)
}

func updateExercise(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)

    err := DataStore.Exercises.Update(id, r)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusNotFound)
    }
    
    w.WriteHeader(http.StatusOK)
}

func deleteExercise(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)

    err := DataStore.Exercises.Delete(id)
    if err != nil {
        log.Println(err)
    }

    w.WriteHeader(http.StatusOK)
}
