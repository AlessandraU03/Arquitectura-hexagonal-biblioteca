package dependencies

import (
	"demo/src/internal/books/application"
	"demo/src/internal/books/domain/entities"
)

type ListBooksController struct {
	useCase application.ListBooks
}

func NewListBooksController(useCase application.ListBooks) *ListBooksController {
	return &ListBooksController{useCase: useCase}
}

func (lp_c *ListBooksController) Execute() ([]*entities.Book, error) {
	return lp_c.useCase.Execute()
}
