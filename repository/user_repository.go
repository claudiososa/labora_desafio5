package repository

import (
	"context"

	"github.com/claudiososa/labora_desafio5/models" // Importa el paquete de modelos
)

// ItemRepository representa un repositorio de items.
type UserRepository struct {
	// Aquí podrías incluir cualquier dependencia necesaria, como una conexión a la base de datos.
}

// Save guarda un nuevo ítem en el repositorio.
func (r *UserRepository) Save(ctx context.Context, item *models.User) error {
	// Aquí implementas la lógica para guardar el ítem en la base de datos.
	return nil
}

func (r *UserRepository) Edit(ctx context.Context, item *models.User) error {
	// Aquí implementas la lógica para guardar el ítem en la base de datos.
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, item *models.User) error {
	// Aquí implementas la lógica para guardar el ítem en la base de datos.
	return nil
}
