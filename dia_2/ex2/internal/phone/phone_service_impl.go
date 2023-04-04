package phone

import (
	"github.com/PauloVML/GoWeb/dia_2/ex2/internal/domain"
)

type TelefonoServiceImpl struct {
	Storage Repository
}

//Primero debemos implementar las funciones declaradas en el service (interface)

func (tsi *TelefonoServiceImpl) Save(telefono *domain.Telefono) error {

	//Creamos el producto en la base de datos
	//Y gestionamos los errores
	var err error
	if err = tsi.Storage.Save(telefono); err != nil {
		//errors.Is(err, unErrorCualquiera){
		//	entonces tal cosa
		//}
	}
	return nil
}

func (tsi *TelefonoServiceImpl) GetAll() ([]domain.Telefono, error) {
	telefonos, err := tsi.Storage.GetAll()
	if err != nil {
		//manejo de errores
	}
	//En este caso que DEVUELVE telefonos no usamos un puntero, si no que este
	//lo usamos cuando lo estamos por modificar, es decir cuando viene por parametro
	return telefonos, nil
}
