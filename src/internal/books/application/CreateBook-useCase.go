package application

import (
	"demo/src/internal/books/domain"
	"demo/src/internal/books/domain/entities"
)

type CreateBook struct {
	db domain.IBook
}

func NewCreateBook(db domain.IBook) *CreateBook{
	return &CreateBook{db: db}
}

func (cp *CreateBook) Execute(book *entities.Book) error{
	return cp.db.Save(book)
}
