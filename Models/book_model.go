package Models

type Book struct {
	id        int
	title     string
	author_id int
}

func (book *Book) GetId() int {
	return book.id
}

func (book *Book) GetTitle() string {
	return book.title
}

func (book *Book) GetAuthor_id() int {
	return book.author_id
}

func (book *Book) SetId(id int) *Book {
	book.id = id
	return book
}

func (book *Book) SetTitle(title string) *Book {
	book.title = title
	return book
}

func (book *Book) SetAuthor_id(author_id int) *Book {
	book.author_id = author_id
	return book
}
