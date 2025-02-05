package application

import (
	"demo/src/internal/comics/domain"
)

type UpdateComic struct {
	db domain.IComic
}

func NewUpdateComic(db domain.IComic) *UpdateComic{
	return &UpdateComic{db: db}
}

func (up *UpdateComic) Execute(id int32 ,name string, autor string, editorial string){
	up.db.Update(id,name, autor, editorial)
}
