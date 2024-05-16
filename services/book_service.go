package services

import (
	"context"

	"github.com/claudiososa/labora_desafio5/models"
	"github.com/claudiososa/labora_desafio5/repository"
)

type BookService struct{}

func (BookService) GetBooksByStatus(status string) ([]models.Book, error) {

	repo := &repository.BookRepository{}

	books, err := repo.GetBooksByStatus(context.Background(), status)

	if err != nil {
		// Maneja el error
		return nil, err
	}
	return books, nil
}

func (BookService) SearchBooksByTitileAndAuthor(titieOrAuthor string) ([]models.Book, error) {

	repo := &repository.BookRepository{}

	books, err := repo.SearchBooksByTitileAndAuthor(context.Background(), titieOrAuthor)

	if err != nil {
		// Maneja el error
		return nil, err
	}
	return books, nil
}
