package entities

type Book struct{
	ID int32 
	Name string
	Autor string
	Categoria string
}

func NewBook(name string, autor string, categoria string) *Book{
	return &Book{
		ID: 0,
		Name: name,
		Autor: autor,
		Categoria: categoria,
	}
}

func (l *Book) GetName() string{
	return l.Name
}

func (l *Book) SetName(name string){
	l.Name = name
}

