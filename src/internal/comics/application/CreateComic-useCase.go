package application

import (
	"demo/src/internal/comics/domain"
	"demo/src/internal/comics/domain/entities"
)

type CreateComic struct {
	db domain.IComic
}

func NewCreateComic(db domain.IComic) *CreateComic{
	return &CreateComic{db: db}
}

func (cp *CreateComic) Execute(comic *entities.Comic) error{
	return cp.db.Save(comic)
}
