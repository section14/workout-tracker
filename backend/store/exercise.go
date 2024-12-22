package store

type Exercise struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (Exercise *e) GetAll() []Exercise {

}

func (Exercise *e) Get(id int) Exercise {

}

func (Exercise *e) Save() Exercise {

}

func (Exercise *e) Update(id int) err {
    return nil
}

func (Exercise *e) Delete(id int) err {
    return nil
}
