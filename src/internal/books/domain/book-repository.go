package domain

type IBook interface {
	Save(name string, autor string, categoria string) error
	GetAll() ([]map[string]interface{}, error)
	GetById(int32) (map[string]interface{}, error)
	Update(id int32, name string, autor string, categoria string) error
	Delete(id int32) error
}
