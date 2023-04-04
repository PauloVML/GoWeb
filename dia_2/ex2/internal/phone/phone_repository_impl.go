package phone

import "github.com/PauloVML/GoWeb/dia_2/ex2/internal/domain"

// Esta es la implementación entonces debe compartir los métodos
type RepositoryImpl struct {
	data []domain.Telefono
}

// acá ponemos el puntero para que los cambios los haga sobre el data en memoria
func (repoImpl *RepositoryImpl) GetAll() (resultado []domain.Telefono, err error) {
	resultado = repoImpl.data
	return
}

func (repoImpl *RepositoryImpl) Save(telefono *domain.Telefono) error {
	telefono.Id = len(repoImpl.data) + 1
	//Aca también agregamos un puntero
	repoImpl.data = append(repoImpl.data, *telefono)
	return nil
}
