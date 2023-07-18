package db

import (
	"database/sql"
	"fmt"
	"log"
	types "prevue/pkg/types"

	_ "github.com/lib/pq"
)

func Connect(dbName, dbUser, dbPassword, dbHost, dbPort string) *sql.DB {
	// Connect to the database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database...")
	return db
}

func Close(db *sql.DB) {
	db.Close()
	log.Println("Closed connection to database...")
}

func CreateTables(db *sql.DB) {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Users (user_id serial PRIMARY KEY, user_name VARCHAR ( 100 ) UNIQUE NOT NULL, email VARCHAR ( 100 ) UNIQUE NOT NULL, passwoard VARCHAR ( 100 ) NOT NULL)")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Created table Users")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Models (model_id serial PRIMARY KEY, model_name TEXT NOT NULL, connector_name TEXT NOT NULL, architecture bytea, weights bytea)")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Created table Models")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Projects (project_id serial PRIMARY KEY, project_name TEXT NOT NULL, user_id INT NOT NULL, model_id INT NOT NULL, FOREIGN KEY (user_id) REFERENCES Users (user_id), FOREIGN KEY (model_id) REFERENCES Models (model_id))")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Created table Projects")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Metrics (metric_id serial PRIMARY KEY, model_id INT NOT NULL, epoch INT, batch INT, loss_name TEXT, loss_value FLOAT, metric_name TEXT, metric_value FLOAT, FOREIGN KEY (model_id) REFERENCES Models (model_id))")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Created table Metrics")
}

func SessionData(db *sql.DB, data types.SessionData) (int, error) {
	_, err := db.Exec("INSERT INTO Users (user_name, email, passwoard) SELECT CAST($1 AS VARCHAR),CAST($2 AS VARCHAR), $3 WHERE NOT EXISTS (SELECT user_name, email  FROM Users WHERE user_name = $1 and email = $2)", data.User.UserName, data.User.Email, data.User.Passwoard)
	if err != nil {
		log.Printf("Failed to insert data into user table: %v", err)
		return 0, err
	}
	log.Println("Inserted data to user table.")

	modelId := 0
	err = db.QueryRow("insert into Models (model_name, connector_name, architecture, weights) values ($1, $2, $3, $4) RETURNING model_id;", data.Models.ModelName, data.Models.ConnectorName, data.Models.Architecture, data.Models.Weights).Scan(&modelId)
	if err != nil {
		log.Printf("Failed to insert data into model table: %v", err)
		return 0, err
	}
	log.Println("Inserted data to models table.")

	_, err = db.Exec("insert into Projects (project_name, user_id, model_id) select $1, user_id, $3 from users where user_name = $2;", data.Project.ProjectName, data.User.UserName, modelId)
	if err != nil {
		log.Printf("Failed to insert data into project table: %v", err)
		return 0, err
	}
	log.Println("Inserted data to project table.")

	return modelId, nil

}
