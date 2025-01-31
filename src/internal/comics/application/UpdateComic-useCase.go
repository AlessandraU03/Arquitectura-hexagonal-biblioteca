package application

import (
	"demo/src/internal/comics/domain"
	"demo/src/internal/comics/domain/entities"
)

type UpdateComic struct {
	db domain.IComic
}

func NewUpdateComic(db domain.IComic) *UpdateComic{
	return &UpdateComic{db: db}
}

func (up *UpdateComic) Execute(comic *entities.Comic) error{
	return up.db.Update(comic)
}
