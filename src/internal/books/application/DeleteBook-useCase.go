package application

import "demo/src/internal/books/domain"

type DeleteBook struct {
	db domain.IBook
}

func NewDeleteBook(db domain.IBook) *DeleteBook {
    return &DeleteBook{db: db}
}

func (dp *DeleteBook) Delete(id int32) error {
    return dp.db.Delete(id)
}
