package models

type Author struct {
	Id   int
	Name string
}

func (author *Author) NewAuthor(id int, name string) *Author {
	return &Author{
		Id:   id,
		Name: name,
	}
}

func (author *Author) GetId() int {
	return author.Id
}

func (author *Author) GetName() string {
	return author.Name
}

func (author *Author) SetId(id int) *Author {
	author.Id = id
	return author
}

func (author *Author) SetName(name string) *Author {
	author.Name = name
	return author
}
