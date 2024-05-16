package models

type Author struct {
	id   *int
	Name string
}

func (author *Author) NewAuthor(id *int, name string) *Author {
	return &Author{
		id:   id,
		Name: name,
	}
}

func (author *Author) GetId() *int {
	return author.id
}

func (author *Author) GetName() string {
	return author.Name
}

func (author *Author) SetId(id *int) *Author {
	author.id = id
	return author
}

func (author *Author) SetName(name string) *Author {
	author.Name = name
	return author
}
