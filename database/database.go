package database

import (
	"crud/models"
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	//_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDB() {

	username := "posttest"
	password := "test1234"
	host := "localhost"
	port := "5432"
	databaseName := "crud"

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		username, password, host, port, databaseName)

	var err error

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database!")

	sql := `CREATE TABLE IF NOT EXISTS pets (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        photourls TEXT NOT NULL,
        tags TEXT NOT NULL,
        status INT NOT NULL
    );`

	_, err = db.Exec(sql)
	if err != nil {
		log.Println("Error in Creating the table ", err)
		return
	}

	log.Println("Table created successfully!!")
}

// InsertIntoPet insert a single animal into pet table
func InsertIntoPet(data models.Pet) error {
	sql := `INSERT INTO pets ("name","photourls","tags","status")
    VALUES($1,$2,$3,$4)`
	_, err := db.Exec(sql, data.Name, pq.Array(data.PhotoUrls), pq.Array(data.Tags), data.Status)
	if err != nil {
		return err
	}
	return nil
}

// GetPetByID will get pet by ID
func GetPetByID(id int) (models.Pet, error) {
	sql := "SELECT * FROM pets WHERE id = $1;"

	var data models.Pet

	err := db.QueryRow(sql, id).Scan(&data.Id, &data.Name, pq.Array(&data.PhotoUrls), pq.Array(&data.Tags), &data.Status)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetPetByStatus function to get pet by status
func GetPetByStatus(status int) ([]models.Pet, error) {
	sql := "SELECT * FROM pets WHERE status = $1;"

	var data []models.Pet

	rows, err := db.Query(sql, status)
	if err != nil {
		return data, err
	}

	defer rows.Close()

	var id, statuss int
	var name string
	var photo_urls, tags []string

	for rows.Next() {

		err := rows.Scan(&id, &name, pq.Array(&photo_urls), pq.Array(&tags), &statuss)
		if err != nil {
			return data, err
		}

		d := models.Pet{
			Id:        id,
			Name:      name,
			PhotoUrls: photo_urls,
			Tags:      tags,
			Status:    statuss,
		}

		data = append(data, d)

	}

	return data, nil

}

// DeletePetByStatus delete the Pet by ID of the Pet
func DeletePetByStatus(id int) error {
	sql := "DELETE from pets WHERE id=$1;"

	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}

// GetCountOfID check whether the Pet with particular id is  present or not
func GetCountOfID(id int) (int, error) {
	var count int

	sql := "SELECT count(name) from pets WHERE id=$1;"
	err := db.QueryRow(sql, id).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// FullUpdateByID update the whole data for particular id
func FullUpdateByID(id int, data models.Pet) error {
	sql := `UPDATE pets set name=$1,photourls=$2,tags=$3,status=$4
            WHERE id=$5;`

	_, err := db.Exec(sql, data.Name, pq.Array(data.PhotoUrls), pq.Array(data.Tags), data.Status, id)
	if err != nil {
		return err
	}
	return nil
}
