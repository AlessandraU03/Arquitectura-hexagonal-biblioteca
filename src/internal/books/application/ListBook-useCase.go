package application

import (
	"demo/src/internal/books/domain"
	"demo/src/internal/books/domain/entities"
)

type ListBooks struct {
	db domain.IBook
}

func NewListBook(db domain.IBook) *ListBooks {
	return &ListBooks{db: db}
}

func (lp *ListBooks) Execute() ([]*entities.Book, error){
	return lp.db.GetAll()
}