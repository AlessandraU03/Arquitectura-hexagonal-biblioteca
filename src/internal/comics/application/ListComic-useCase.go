package application

import (
	"demo/src/internal/comics/domain"
	"demo/src/internal/comics/domain/entities"
)

type ListComics struct {
	db domain.IComic
}

func NewListComics(db domain.IComic) *ListComics{
	return &ListComics{db: db}
}

func (lc *ListComics) Execute() ([]*entities.Comic, error){
	return lc.db.GetAll()
}

