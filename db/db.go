package db

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/pressly/goose/v3"
)

func InitDb() *sql.DB {
    src := "root:@tcp(127.0.0.1:3306)/three_layer?parseTime=true"
    db, err := sql.Open("mysql", src)
    if err != nil {
        log.Fatal("Failed to open database:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("✅ Connected to MySQL Database: three_layer")

    // Run migrations automatically
    if err := runMigrations(db); err != nil {
        log.Fatal("Failed to run migrations:", err)
    }

    return db
}

func runMigrations(db *sql.DB) error {
    if err := goose.SetDialect("mysql"); err != nil {
        return err
    }

    if err := goose.Up(db, "./migrations"); err != nil {
        return err
    }

    log.Println("✅ Migrations applied successfully!")
    return nil
}
