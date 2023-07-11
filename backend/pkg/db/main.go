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
	// Create table if it doesn't exist
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS records (uid TEXT, project_name TEXT, connector_name TEXT, accuracy REAL, mse REAL)")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Created table...")
}

func Insert(db *sql.DB, data types.ConnectorData) {
	// Insert data into table
	_, err := db.Exec("INSERT INTO records (uid, project_name, connector_name, accuracy, mse) VALUES ($1, $2, $3, $4, $5)", data.Uid, data.ProjectName, data.ConnectorName, data.Metrics.Accuracy, data.Metrics.MSE)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
	log.Println("Inserted data...")
}
