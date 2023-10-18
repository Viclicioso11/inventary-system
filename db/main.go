package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	fmt.Println(dbUsername)
	fmt.Println(dbPassword)

	connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5435/go-test-database?sslmode=disable", dbUsername, dbPassword)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		fmt.Println(err)
	}

	migrationPath, _ := filepath.Abs("./migrations")

	fmt.Println(migrationPath)

	migrationPath = strings.Replace(migrationPath, "C:\\", "file:///", 1)

	fmt.Println(migrationPath)

	m, err := migrate.NewWithDatabaseInstance(
		//"file:///Users\\USER\\Documents\\InventorySystem\\db\\migrations", this is the correct path (absolute)
		migrationPath,
		"go-test-database", driver)

	if err != nil {
		panic(err)
	}
	if m != nil {
		err := m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

		if err != nil {
			panic(err)
		}
	}
}
