package application

import (
	"demo/src/internal/books/domain"
)

type CreateBook struct {
	db domain.IBook
}

func NewCreateBook(db domain.IBook) *CreateBook{
	return &CreateBook{db: db}
}

func (cp *CreateBook) Execute(name string, autor string, categoria string){
	cp.db.Save(name, autor, categoria)
}
