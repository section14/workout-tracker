package store

type Workout struct {
	Id        int64     `json:"id"`
	Exercises []Action  `json:"exercises"`
	Date      time.Time `json:"date_modified"`
}

func (Workout *w) GetAll() []Workout {

}

func (Workout *w) Get(id int) Workout {

}

func (Workout *w) Update(id int) err {
	return nil
}

func (Workout *w) Delete(id int) err {
	return nil
}
