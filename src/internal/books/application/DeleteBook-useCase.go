package application

import "demo/src/internal/books/domain"

type DeleteBook struct {
	db domain.IBook
}

func NewDeleteBook(db domain.IBook) *DeleteBook {
    return &DeleteBook{db: db}
}

func (dp *DeleteBook) Execute(id int32) {
    dp.db.Delete(id)
}
