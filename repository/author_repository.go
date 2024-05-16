package repository

import (
	"context"

	"github.com/claudiososa/labora_desafio5/models"
	"github.com/claudiososa/labora_desafio5/utils"
)

type AuthorRepository struct{}

func (r *AuthorRepository) Save(ctx context.Context, author *models.Author) error {
	query := "INSERT INTO authors (name) VALUES (?)"

	_, err := utils.DB.Exec(query, author.Name)
	if err != nil {
		return err
	}
	return nil

}

func (r *AuthorRepository) Edit(ctx context.Context, item *models.Author) error {
	return nil
}

func (r *AuthorRepository) Delete(ctx context.Context, item *models.Author) error {
	return nil
}
