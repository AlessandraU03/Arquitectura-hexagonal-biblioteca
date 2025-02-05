package entities

type Book struct{
	id int32 
	name string
	autor string
	categoria string
}

func NewBook(name string, autor string, categoria string) *Book{
	return &Book{
		id: 0,
		name: name,
		autor: autor,
		categoria: categoria,
	}
}


func (l *Book) GetName() string{
	return l.name
}

func (l *Book) SetName(name string){
	l.name = name
}

