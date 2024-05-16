package services

import (
	"context"
	"fmt"

	"github.com/claudiososa/labora_desafio5/models"
	"github.com/claudiososa/labora_desafio5/repository"
)

type AuthorService struct{}

func (AuthorService) Create(name string) error {

	var author models.Author
	author1 := author.NewAuthor(0, name)

	repo := &repository.AuthorRepository{}

	err := repo.Save(context.Background(), author1)

	if err != nil {
		// Maneja el error
		return fmt.Errorf("Error al crear el autor: %w", err)
	}
	return nil
}
