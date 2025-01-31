package database

import (
	"demo/src/core"
	"demo/src/internal/books/domain/entities"
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



func (mysql *MySQLBooks) Save(Book *entities.Book) error {
    query := "INSERT INTO books (name, autor, categoria) VALUES (?, ?, ?)" // Añadimos categoria aquí
    result, err := mysql.conn.ExecutePreparedQuery(query, Book.Name, Book.Autor, Book.Categoria)
    if err != nil {
        return fmt.Errorf("error al guardar el libro: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
    return nil
}


// GetAll - Obtiene todos los Bookos
// GetAll - Obtiene todos los Bookos
func (mysql *MySQLBooks) GetAll() ([]*entities.Book, error) {
    query := "SELECT id, name, autor, categoria FROM books" // Añadimos categoria aquí
    rows := mysql.conn.FetchRows(query)
    defer rows.Close()

    var Books []*entities.Book
    for rows.Next() {
        Book := &entities.Book{}
        if err := rows.Scan(&Book.ID, &Book.Name, &Book.Autor, &Book.Categoria); err != nil {
            return nil, fmt.Errorf("error al escanear el libro: %w", err)
        }
        Books = append(Books, Book)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterando sobre los Bookos: %w", err)
    }

    return Books, nil
}

// GetByID - Obtiene un Booko por ID
func (mysql *MySQLBooks) GetByID(id int32) (*entities.Book, error) {
    query := "SELECT id, name, autor, categoria FROM books WHERE id = ?"
    rows := mysql.conn.FetchRows(query, id)
    defer rows.Close()

    if rows.Next() {
        Book := &entities.Book{}
        if err := rows.Scan(&Book.ID, &Book.Name, &Book.Autor, &Book.Categoria); err != nil {
            return nil, fmt.Errorf("error al escanear el Booko: %w", err)
        }
        return Book, nil
    }

    return nil, fmt.Errorf("Book con ID %d no encontrado", id)
}


// Update - Actualiza un Booko
func (mysql *MySQLBooks) Update(Book *entities.Book) error {
	query := "UPDATE books SET name = ?, autor = ?, categoria = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, Book.Name, Book.Autor, Book.Categoria, Book.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el Booko: %w", err)
	}
	return nil
}

// Delete - Elimina un Booko
func (mysql *MySQLBooks) Delete(id int32) error {
    query := "DELETE FROM books WHERE id = ?"
    _, err := mysql.conn.ExecutePreparedQuery(query, id)
    if err != nil {
        return fmt.Errorf("error al eliminar el Booko: %w", err)
    }
    return nil
}