package store

type Action struct {
	Id       int64    `json:"id"`
	Exercise Exercise `json:"exercise"`
	Sets     int64    `json:"sets"`
	Reps     int64    `json:"reps"`
}

func (a* Action) GetAll() []Action {
    return []Action{}
}

func (a* Action) Get(id int) Action {
    return Action{}
}

func (a* Action) Update(id int) error {
	return nil
}

func (a* Action) Delete(id int) error {
	return nil
}
