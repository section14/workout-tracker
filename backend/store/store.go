package store

import "time"

type Store struct {
	Exercises []Exercise `json:"exercises"`
	Workouts  []Workout  `json:"workouts"`
	Updated   time.Time  `json:"date_modified"`
}

func (Store *s) Get() Store {

}

func (Store *s) Update(id int) err {
	return nil
}

func (Store *s) Delete(id int) err {
	return nil
}
