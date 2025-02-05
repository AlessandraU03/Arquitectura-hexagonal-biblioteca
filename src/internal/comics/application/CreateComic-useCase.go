package application

import (
	"demo/src/internal/comics/domain"
)

type CreateComic struct {
	db domain.IComic
}

func NewCreateComic(db domain.IComic) *CreateComic{
	return &CreateComic{db: db}
}

func (cp *CreateComic) Execute(name string, autor string, editorial string){
	cp.db.Save(name, autor, editorial)	
}
