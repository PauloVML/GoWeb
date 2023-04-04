package phone

import "github.com/PauloVML/GoWeb/dia_2/ex2/internal/domain"

type TelefonoService interface {
	Save(*domain.Telefono) error
	GetAll() ([]domain.Telefono, error)
}
