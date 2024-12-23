package store

import (
	"net/http"
	"time"
)

type Store struct {
	Exercises Exercises `json:"storeExercises"`
	Workouts  Workouts `json:"storeWorkouts"`
	Updated   time.Time `json:"updated"`
}

type Crud interface {
	GetAll() ([]byte, error)
	Get() ([]byte, error)
	Create(r *http.Request) error
	Update(int64, r *http.Request) error
	Delete(int64) error
}

func (s *Store) Get() Store {
	return Store{}
}

func (s *Store) Update(id int) error {
	return nil
}

func (s *Store) Delete(id int) error {
	return nil
}
