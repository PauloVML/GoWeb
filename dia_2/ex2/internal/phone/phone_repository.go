package phone

import (
	"errors"
	"github.com/PauloVML/GoWeb/dia_2/ex2/internal/domain"
)

// acá entiendo que podemos declarar algunos errores
var (
	ErrElProductoYaExiste = errors.New("Este producto ya existe")
)

type Repository interface {
	//acá como llega algo por parametro y no es algo que yo devuelvo, utilizamos un puntero
	Save(*domain.Telefono) error
	GetAll() ([]domain.Telefono, error)
}
