package controllers

import (
	"demo/src/internal/books/application"
	"demo/src/internal/books/domain/entities"
)

type CreateBookController struct {
	useCase application.CreateBook
}

func NewCreateBookController(useCase application.CreateBook) *CreateBookController {
	return &CreateBookController{useCase: useCase}
}

func (cp_c *CreateBookController) Execute(book *entities.Book) error {
	return cp_c.useCase.Execute(book)
}
