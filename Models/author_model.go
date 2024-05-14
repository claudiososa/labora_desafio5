package Models

type Author struct {
	id   int
	name string
}

func (author *Author) GetId() int {
	return author.id
}

func (author *Author) GetName() string {
	return author.name
}

func (author *Author) SetId(id int) *Author {
	author.id = id
	return author
}

func (author *Author) SetName(name string) *Author {
	author.name = name
	return author
}
