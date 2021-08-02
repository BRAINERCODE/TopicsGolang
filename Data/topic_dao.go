package data

import (
	"Topic/Database"
	"Topic/Models"
	"database/sql"
	"fmt"
)

type Implement struct {
	db *sql.DB
}

var t Models.Topic
var ArrayT []Models.Topic

func (i Implement) GetAllTopics() ([]Models.Topic, error) {
	db, err := Database.GetMySqlClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, _ := db.Query("SELECT * FROM topics")

	if err != nil {
		return nil, err
	}
	defer filas.Close()

	for filas.Next() {
		err = filas.Scan(&t.ID, &t.Categoria, &t.Lenguaje, &t.Estado, &t.Descripcion)
		if err != nil {
			return nil, err
		}
		ArrayT = append(ArrayT, t)
	}
	return ArrayT, nil

}

func (i Implement) GetTopicById(id int64) (Models.Topic, error) {

	db, err := Database.GetMySqlClient()
	if err != nil {
		return Models.Topic{}, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT * FROM topics WHERE IdTopic = ?", id)

	if err != nil {
		return Models.Topic{}, err
	}
	defer filas.Close()

	for filas.Next() {

		err = filas.Scan(&t.ID, &t.Categoria, &t.Lenguaje, &t.Estado, &t.Descripcion)
		if err != nil {

			return Models.Topic{}, err
		}

	}
	return t, nil

}
func (i Implement) PostTopic(c Models.Topic) error {

	db, err := Database.GetMySqlClient()
	if err != nil {
		return err
	}
	defer db.Close()

	sentencia, err := db.Prepare("INSERT INTO topics (IdTopic,Categoria,Lenguaje,Estado,Descripcion) VALUES(?,?, ?, ?,?)")
	if err != nil {
		return err
	}
	defer sentencia.Close()
	_, err = sentencia.Exec(c.ID, c.Categoria, c.Lenguaje, c.Estado, c.Descripcion)

	if err != nil {
		return err
	}

	return nil

}
func (i Implement) UpdateTopic(c Models.Topic, id int) error {
	db, err := Database.GetMySqlClient()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE topics SET Categoria = ?, Lenguaje = ?, Estado = ?, Descripcion = ? WHERE IdTopic = ?")
	fmt.Println("Sentencia", sentenciaPreparada)
	if err != nil {
		return err
	}

	defer sentenciaPreparada.Close()
	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaPreparada.Exec(c.Categoria, c.Lenguaje, c.Estado, c.Descripcion, id)
	return err // Ya sea nil o sea un error, lo manejaremos desde donde hacemos la llamada
}
func (i Implement) DeleteTopic(id int) error {
	db, err := Database.GetMySqlClient()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM topics WHERE IdTopic = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
