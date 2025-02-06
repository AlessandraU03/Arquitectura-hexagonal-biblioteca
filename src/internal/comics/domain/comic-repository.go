package domain

type IComic interface {
	Save(name string, autor string, editorial string) error
	GetAll() ([]map[string]interface{}, error)
	GetById(int32) (map[string]interface{}, error)
	Update(id int32, name string, autor string, editorial string) error
	Delete(id int32) error
}
