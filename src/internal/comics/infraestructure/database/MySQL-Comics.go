package database

import (
	"demo/src/core"
	"demo/src/internal/comics/domain/entities"
	"fmt"
	"log"
)

type MySQLComics struct {
	conn *core.Conn_MySQL
}

func NewMySQLComics() *MySQLComics {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQLComics{conn: conn}
}



func (mysql *MySQLComics) Save(Comic *entities.Comic) error {
    query := "INSERT INTO Comics (name, autor, Editorial) VALUES (?, ?, ?)" // Añadimos Editorial aquí
    result, err := mysql.conn.ExecutePreparedQuery(query, Comic.Name, Comic.Autor, Comic.Editorial)
    if err != nil {
        return fmt.Errorf("error al guardar el libro: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
    return nil
}


// GetAll - Obtiene todos los Comicos
// GetAll - Obtiene todos los Comicos
func (mysql *MySQLComics) GetAll() ([]*entities.Comic, error) {
    query := "SELECT id, name, autor, editorial FROM Comics" // Añadimos Editorial aquí
    rows := mysql.conn.FetchRows(query)
    defer rows.Close()

    var Comics []*entities.Comic
    for rows.Next() {
        Comic := &entities.Comic{}
        if err := rows.Scan(&Comic.Id, &Comic.Name, &Comic.Autor, &Comic.Editorial); err != nil {
            return nil, fmt.Errorf("error al escanear el libro: %w", err)
        }
        Comics = append(Comics, Comic)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterando sobre los Comicos: %w", err)
    }

    return Comics, nil
}

// GetByID - Obtiene un Comico por ID
func (mysql *MySQLComics) GetByID(id int32) (*entities.Comic, error) {
    query := "SELECT id, name, autor, Editorial FROM Comics WHERE id = ?"
    rows := mysql.conn.FetchRows(query, id)
    defer rows.Close()

    if rows.Next() {
        Comic := &entities.Comic{}
        if err := rows.Scan(&Comic.Id, &Comic.Name, &Comic.Autor, &Comic.Editorial); err != nil {
            return nil, fmt.Errorf("error al escanear el Comico: %w", err)
        }
        return Comic, nil
    }

    return nil, fmt.Errorf("Comic con ID %d no encontrado", id)
}


// Update - Actualiza un Comico
func (mysql *MySQLComics) Update(Comic *entities.Comic) error {
	query := "UPDATE Comics SET name = ?, autor = ?, Editorial = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, Comic.Name, Comic.Autor, Comic.Editorial, Comic.Id)
	if err != nil {
		return fmt.Errorf("error al actualizar el Comico: %w", err)
	}
	return nil
}

// Delete - Elimina un Comico
func (mysql *MySQLComics) Delete(id int32) error {
    query := "DELETE FROM Comics WHERE id = ?"
    _, err := mysql.conn.ExecutePreparedQuery(query, id)
    if err != nil {
        return fmt.Errorf("error al eliminar el Comico: %w", err)
    }
    return nil
}