package application

import (
	"demo/src/internal/books/domain"
)

type ListBooks struct {
	db domain.IBook
}

func NewListBook(db domain.IBook) *ListBooks {
	return &ListBooks{db: db}
}

func (lp *ListBooks) Execute() ([]map[string]interface{}, error){
	return lp.db.GetAll()
}