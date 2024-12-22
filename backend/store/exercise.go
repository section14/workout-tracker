package store

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
)

type Exercise struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Exercises struct {
    Exercises []Exercise `json:"exercises"`
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

func (e *Exercises) GetAll() ([]byte, error) {
	var empty []byte

	jsonData, err := json.Marshal(e.Exercises)
	if err != nil {
		return empty, errors.New("couldn't get all exercises")
	}

	return jsonData, nil
}

func (e *Exercises) Get(id int64) ([]byte, error) {
	var empty []byte
	var ex Exercise

	for _, es := range e.Exercises {
		if id == es.Id {
			ex = es
		}
	}

	jsonData, err := json.Marshal(ex)
	if err != nil {
		return empty, errors.New("couldn't get exercise")
	}

	return jsonData, nil
}

func (e *Exercises) Create(r *http.Request) error {
	length := len(e.Exercises)
	newId := e.Exercises[length-1].Id + 1

	type Temp struct {
		Name string `json:"name"`
	}

	var t Temp

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return err
	}

	ex := Exercise{Id: newId, Name: t.Name}
	e.Exercises = append(e.Exercises, ex)

	return nil
}

func (e *Exercises) Update(id int64, r *http.Request) error {
	var idx int
	var ex Exercise

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ex)
	if err != nil {
		return err
	}

	for i, es := range e.Exercises {
		if id == es.Id {
			idx = i
		}
	}

    e.Exercises = slices.Replace(e.Exercises, idx, idx+1, ex)

	return nil
}

func (e *Exercises) Delete(id int64) error {
    var idx int

	for i, es := range e.Exercises {
		if id == es.Id {
			idx = i
		}
	}

    e.Exercises = slices.Delete(e.Exercises, idx, idx+1)

    return nil
}
