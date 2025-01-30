package controllers

import (
	"demo/src/internal/books/application"
	"demo/src/internal/books/domain/entities"
)

type ListBookController struct {
	useCase application.ListBooks
}

func NewListBookController(useCase application.ListBooks) *ListBookController {
	return &ListBookController{useCase: useCase}
}

func (lp_c *ListBookController) Execute() ([]*entities.Book, error) {
	return lp_c.useCase.Execute()
}
