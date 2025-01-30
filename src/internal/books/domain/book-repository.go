package domain

import "demo/src/internal/books/domain/entities"

type IBook interface {
	Save(Book *entities.Book) error
	GetAll() ([]*entities.Book, error)
	GetByID(id int32) (*entities.Book, error)
	Update(Book *entities.Book) error
	Delete(id int32) error
}
