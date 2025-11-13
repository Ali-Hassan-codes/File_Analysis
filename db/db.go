package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	src := "root:@tcp(127.0.0.1:3306)/three_layer?parseTime=true"
	db, err := sql.Open("mysql", src)
	if err != nil {
		log.Fatal("Failed to Open database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("✅ Connected to MySQL Database: file_analyzer_2")

	// Create users table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username varchar(255) NOT NULL,
			email varchar(255) NOT NULL UNIQUE,
			password varchar(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	// Create file_info table for File Analyzer
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS file_info (
        id INT AUTO_INCREMENT PRIMARY KEY,
        file_name VARCHAR(255),
        paragraphs INT,
        line_count INT,
        word_count INT,
        char_count INT,
        alphabetic INT,
        digits INT,        -- renamed from numeric
        vowels INT,
        non_vowels INT
    );
`)
if err != nil {
    log.Fatal("Failed to create file_info table:", err)
}

	return db
}
