package application

import (
	"demo/src/internal/comics/domain"
)

type ListComics struct {
	db domain.IComic
}

func NewListComics(db domain.IComic) *ListComics{
	return &ListComics{db: db}
}

func (lc *ListComics) Execute() ([]map[string]interface{}, error){
	return lc.db.GetAll()
}

