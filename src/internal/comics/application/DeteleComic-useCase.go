package application

import (
	"demo/src/internal/comics/domain"
)

type DeleteComic struct {
	db domain.IComic
}

func NewDeleteComic(db domain.IComic) *DeleteComic{
	return &DeleteComic{db: db}
}

func (cp *DeleteComic) Execute(id int32){
	 cp.db.Delete(id)
}
