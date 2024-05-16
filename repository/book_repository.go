package repository

import (
	"context"
	"fmt"

	"github.com/claudiososa/labora_desafio5/models"
	"github.com/claudiososa/labora_desafio5/utils"
)

type BookRepository struct{}

func (r *BookRepository) Save(ctx context.Context, book *models.Book) error {
	query := "INSERT INTO books (name) VALUES (?)"

	_, err := utils.DB.Exec(query, book.Title)
	if err != nil {
		return err
	}
	return nil

}

func (r *BookRepository) Edit(ctx context.Context, item *models.Author) error {
	return nil
}

func (r *BookRepository) Delete(ctx context.Context, item *models.Author) error {
	return nil
}

func (r *BookRepository) SearchBooksByTitileAndAuthor(ctx context.Context, titleOrAuthor string) ([]models.Book, error) {
	query := fmt.Sprintf("SELECT b.id, b.title, b.status, a.id, a.name FROM books b LEFT JOIN authors a ON a.id = b.author_id WHERE b.title LIKE '%%%s%%' OR a.name LIKE '%%%s%%'", titleOrAuthor, titleOrAuthor)

	rows, err := utils.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		var author models.Author

		if err := rows.Scan(&book.Id, &book.Title, &book.Status, &author.Id, &author.Name); err != nil {
			return nil, err
		}

		book.Author = author
		//fmt.Printf("Fetched book: with author: %+v\n", book)

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) GetBooksByStatus(ctx context.Context, status string) ([]models.Book, error) {
	query := "SELECT b.id, b.title, b.status, a.id, a.name FROM books b	LEFT JOIN authors a ON a.id = b.author_id WHERE b.status = ?"

	rows, err := utils.DB.QueryContext(ctx, query, status)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		var author models.Author

		if err := rows.Scan(&book.Id, &book.Title, &book.Status, &author.Id, &author.Name); err != nil {
			return nil, err
		}

		book.Author = author
		//fmt.Printf("Fetched book: with author: %+v\n", book)

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
