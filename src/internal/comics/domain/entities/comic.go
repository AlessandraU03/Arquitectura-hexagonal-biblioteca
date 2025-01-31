package entities

type Comic struct{
	Id int32 
	Name string
	Autor string
	Editorial string
}

func NewComic(name string, autor string, editorial string) *Comic{
	return &Comic{
		Id: 0,
		Name: name,
		Autor: autor,
		Editorial: editorial,
	}
}

func (l *Comic) GetName() string{
	return l.Name
}

func (l *Comic) SetName(name string){
	l.Name = name
}



