package database

import (
	"demo/src/core"
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

func (mysql *MySQLComics) Save(name string, autor string, editorial string) error {
    query := "INSERT INTO Comics (name, autor, editorial) VALUES (?, ?, ?)"
    result, err := mysql.conn.ExecutePreparedQuery(query, name, autor, editorial)
    if err != nil {
        return fmt.Errorf("error al guardar el cómic: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
    return nil
}

func (mysql *MySQLComics) GetAll() ([]map[string]interface{}, error) {
    query := "SELECT id, name, autor, editorial FROM Comics"
    rows := mysql.conn.FetchRows(query)
    defer rows.Close()

    var comics []map[string]interface{}
    for rows.Next() {
        var id int32
        var name, autor, editorial string
        if err := rows.Scan(&id, &name, &autor, &editorial); err != nil {
            return nil, fmt.Errorf("error al escanear el cómic: %w", err)
        }
        comic := map[string]interface{}{
            "id":       id,
            "name":     name,
            "autor":    autor,
            "editorial": editorial,
        }
        comics = append(comics, comic)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterando sobre los cómics: %w", err)
    }

    return comics, nil
}

func (mysql *MySQLComics) Update(id int32, name string, autor string, editorial string) error {
	query := "UPDATE Comics SET name = ?, autor = ?, editorial = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, name, autor, editorial, id)
	if err != nil {
		return fmt.Errorf("error al actualizar el cómic: %w", err)
	}
	return nil
}

func (mysql *MySQLComics) Delete(id int32) error {
    query := "DELETE FROM Comics WHERE id = ?"
    _, err := mysql.conn.ExecutePreparedQuery(query, id)
    if err != nil {
        return fmt.Errorf("error al eliminar el cómic: %w", err)
    }
    return nil
}