package store

type Action struct {
	Id       int64    `json:"id"`
	Exercise Exercise `json:"exercise"`
	Sets     int64    `json:"sets"`
	Reps     int64    `json:"reps"`
}

func (Action *a) GetAll() []Action {

}

func (Action *a) Get(id int) Action {

}

func (Action *a) Update(id int) err {
	return nil
}

func (Action *a) Delete(id int) err {
	return nil
}
