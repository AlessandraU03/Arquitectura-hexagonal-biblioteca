package database

import (
	"demo/src/core"
	"fmt"
	"log"
)

type MySQLBooks struct {
	conn *core.Conn_MySQL
}

func NewMySQLBooks() *MySQLBooks {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLBooks{conn: conn}
}

func (mysql *MySQLBooks) Save(name string, autor string, categoria string) error {
	query := "INSERT INTO books (name, autor, categoria) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, autor, categoria)
	if err != nil {
		return fmt.Errorf("error al guardar el libro: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	return nil
}

func (mysql *MySQLBooks) GetAll() ([]map[string]interface{}, error) {
	query := "SELECT id, name, autor, categoria FROM books"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var books []map[string]interface{}
	for rows.Next() {
		var id int32
		var name, autor, categoria string
		if err := rows.Scan(&id, &name, &autor, &categoria); err != nil {
			return nil, fmt.Errorf("error al escanear el libro: %w", err)
		}
		book := map[string]interface{}{
			"id":       id,
			"name":     name,
			"autor":    autor,
			"categoria": categoria,
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre los libros: %w", err)
	}

	return books, nil
}

func (mysql *MySQLBooks) Update(id int32, name string, autor string, categoria string) error {
	query := "UPDATE books SET name = ?, autor = ?, categoria = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, name, autor, categoria, id)
	if err != nil {
		return fmt.Errorf("error al actualizar el libro: %w", err)
	}
	return nil
}

func (mysql *MySQLBooks) Delete(id int32) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el libro: %w", err)
	}
	return nil
}
