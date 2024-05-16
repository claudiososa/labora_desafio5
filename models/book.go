package models

type Book struct {
	Id     int
	Title  string
	Status string
	Author Author
}

func (book *Book) GetId() int {
	return book.Id
}

func (book *Book) GetTitle() string {
	return book.Title
}

func (book *Book) GetAuthor() Author {
	return book.Author
}

func (book *Book) GetStatus() string {
	return book.Status
}

func (book *Book) SetId(Id int) *Book {
	book.Id = Id
	return book
}

func (book *Book) SetTitle(Title string) *Book {
	book.Title = Title
	return book
}

func (book *Book) SetAuthor(Author Author) *Book {
	book.Author = Author
	return book
}

func (book *Book) SetStatus(Status string) *Book {
	book.Status = Status
	return book
}
