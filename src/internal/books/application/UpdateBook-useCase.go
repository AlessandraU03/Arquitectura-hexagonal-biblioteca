package application

import (
	"demo/src/internal/books/domain"
)

type UpdateBook struct {
	db domain.IBook
}

func NewUpdateBook(db domain.IBook) *UpdateBook {
	return &UpdateBook{db: db}
}

func (up *UpdateBook) Execute(id int32, name string, autor string, categoria string) {
	up.db.Update(id, name, autor, categoria)
}