package store

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"time"
)

type Workout struct {
	Id      int64     `json:"id"`
	Actions Actions   `json:"workoutActions"`
	Date    time.Time `json:"date"`
}

type Workouts struct {
	Workouts []Workout `json:"workouts"`
}

func (w *Workouts) GetAll() ([]byte, error) {
	var empty []byte
	jsonData, err := json.Marshal(w.Workouts)
	if err != nil {
		return empty, errors.New("couldn't get all workouts")
	}

	return jsonData, nil
}

func (w *Workouts) Get(id int64) ([]byte, error) {
	var empty []byte
	var wo Workout

	for _, ws := range w.Workouts {
		if id == ws.Id {
			wo = ws
		}
	}

	jsonData, err := json.Marshal(wo)
	if err != nil {
		return empty, errors.New("couldn't get workout")
	}

	return jsonData, nil
}

func (w *Workouts) Create(r *http.Request) error {
	length := len(w.Workouts)
	newId := w.Workouts[length-1].Id + 1

    /*
	type Temp struct {
		Actions Actions   `json:"workoutActions"`
		Date    time.Time `json:"date"`
	}

	var t Temp

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return err
	}
    */

	wo := Workout{Id: newId, Actions: Actions{}, Date: time.Now()}
	w.Workouts = append(w.Workouts, wo)

	return nil
}

func (w *Workouts) Update(id int64, r *http.Request) error {
	var idx int
	var wo Workout

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wo)
	if err != nil {
		return err
	}

	for i, ws := range w.Workouts {
		if id == ws.Id {
			idx = i
		}
	}

	w.Workouts = slices.Replace(w.Workouts, idx, idx+1, wo)

	return nil
}

func (w *Workouts) Delete(id int64) error {
	var idx int

	for i, ws := range w.Workouts {
		if id == ws.Id {
			idx = i
		}
	}

	w.Workouts = slices.Delete(w.Workouts, idx, idx+1)

	return nil
}
