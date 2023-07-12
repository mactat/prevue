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

func Insert(db *sql.DB, data types.ConnectorData) error {
	// Insert data into table
	// _, err := db.Exec("INSERT INTO on_epoch_end (uid, project_name, connector_name, epoch, loss, accuracy) VALUES ($1, $2, $3, $4, $5, $6)", data.User.UserName, data.Project.ProjectName, data.Models.ConnectorName, data.Metrics.Epoch, data.Metrics.LossValue, data.Metrics.MatricsValue)
	// if err != nil {
	// 	log.Fatalf("Failed to insert data: %v", err)
	// }
	// log.Println("Inserted data...")

	_, err := db.Exec("INSERT INTO Users (user_name, email, passwoard) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT user_name, email  FROM Users WHERE user_name = user_name and email = email)", data.User.UserName, data.User.Email, data.User.Passwoard)
	if err != nil {
		log.Println("Failed to insert data: %v", err)
		return err
	}
	log.Println("Inserted data to user table.")
	return nil
}
