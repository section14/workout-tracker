package store

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
)

type Action struct {
	Id         int64 `json:"id"`
	ExerciseId int64 `json:"exerciseId"`
	Sets       int64 `json:"sets"`
	Reps       int64 `json:"reps"`
}

type Actions struct {
	Actions []Action `json:"actions"`
}

/*
type Crud interface {
	GetAll() ([]byte, error)
	Get() ([]byte, error)
	Create(r *http.Request) error
	Update(int64, r *http.Request) error
	Delete(int64) error
}
*/

func (a *Actions) GetAll() ([]byte, error) {
	var empty []byte

	jsonData, err := json.Marshal(a.Actions)
	if err != nil {
		return empty, errors.New("couldn't get all actions")
	}

	return jsonData, nil
}

func (a *Actions) Get(id int64) ([]byte, error) {
	var empty []byte
	var ac Action

	for _, as := range a.Actions {
		if id == as.Id {
			ac = as
		}
	}

	jsonData, err := json.Marshal(ac)
	if err != nil {
		return empty, errors.New("couldn't get action")
	}

	return jsonData, nil
}

func (a *Actions) Create(r *http.Request) error {
	length := len(a.Actions)
	newId := a.Actions[length-1].Id + 1

    /*
	type Temp struct {
		ExerciseId int64 `json:"exerciseId"`
		Sets       int64 `json:"sets"`
		Reps       int64 `json:"reps"`
	}

	var t Temp

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return err
	}
    */

	ac := Action{Id: newId, ExerciseId: 1, Sets: 0, Reps: 0}
	a.Actions = append(a.Actions, ac)

	return nil
}

func (a *Actions) Update(id int64, r *http.Request) error {
	var idx int
	var ac Action

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ac)
	if err != nil {
		return err
	}

	for i, as := range a.Actions {
		if id == as.Id {
			idx = i
		}
	}

	a.Actions = slices.Replace(a.Actions, idx, idx+1, ac)

	return nil
}

func (a *Actions) Delete(id int64) error {
	var idx int

	for i, as := range a.Actions {
		if id == as.Id {
			idx = i
		}
	}

	a.Actions = slices.Delete(a.Actions, idx, idx+1)

	return nil
}
