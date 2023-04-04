package handlers

import "github.com/PauloVML/GoWeb/dia_2/ex2/internal/domain"

type SaveTelefonoDTO struct {
	Marca string `json:"marca" binding:"required"`
}

func (dto *SaveTelefonoDTO) toDomain() domain.Telefono {
	return domain.Telefono{
		Marca: dto.Marca,
	}

}
