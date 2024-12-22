package store

import "time"

type Workout struct {
	Id        int64     `json:"id"`
	Exercises []Action  `json:"exercises"`
	Date      time.Time `json:"date_modified"`
}

func (w* Workout) GetAll() []Workout {
    return []Workout{}
}

func (w* Workout) Get(id int) Workout {
    return Workout{}
}

func (w* Workout) Update(id int) error {
	return nil
}

func (w* Workout) Delete(id int) error {
	return nil
}
