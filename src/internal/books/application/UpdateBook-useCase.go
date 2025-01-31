package application

import (
	"demo/src/internal/books/domain"
	"demo/src/internal/books/domain/entities"
)

type UpdateBook struct {
	db domain.IBook
}

func NewUpdateBook(db domain.IBook) *UpdateBook {
	return &UpdateBook{db: db}
}

func (up *UpdateBook) Execute(book *entities.Book) error {
	return up.db.Update(book)
}