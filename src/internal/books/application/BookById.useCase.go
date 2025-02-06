package application

import (
	"demo/src/internal/books/domain"
)

type BookById struct {
	db domain.IBook
}

func NewBookById(db domain.IBook) *BookById {
	return &BookById{db: db}
}

func (lp *BookById) Execute(id int32) (map[string]interface{}, error) {
    return lp.db.GetById(id)
}

