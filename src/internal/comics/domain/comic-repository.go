package domain

import "demo/src/internal/comics/domain/entities"

type IComic interface {
	Save(Comic *entities.Comic) error
	GetAll() ([]*entities.Comic, error)
	GetByID(id int32) (*entities.Comic, error)
	Update(Comic *entities.Comic) error
	Delete(id int32) error
}
