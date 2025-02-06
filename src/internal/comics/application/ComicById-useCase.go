package application

import (
	"demo/src/internal/comics/domain"
)

type ComicById struct {
	db domain.IComic
}

func NewComicById(db domain.IComic) *ComicById {
	return &ComicById{db: db}
}

func (lp *ComicById) Execute(id int32) (map[string]interface{}, error) {
    return lp.db.GetById(id)
}

