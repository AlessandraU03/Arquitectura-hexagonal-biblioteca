package entities

type Comic struct{
	id int32 
	name string
	autor string
	editorial string
}

func NewComic(name string, autor string, editorial string) *Comic{
	return &Comic{
		id: 0,
		name: name,
		autor: autor,
		editorial: editorial,
	}
}

func (l *Comic) GetName() string{
	return l.name
}

func (l *Comic) SetName(name string){
	l.name = name
}



