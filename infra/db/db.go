package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Storage struct {
	db *sql.DB
}

var (
	DB = Storage{}
)

func (s *Storage) Open() (*sql.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USERNAME")
	dbPass := os.Getenv("MYSQL_PASSWORD")

	mysqlPort, _ := strconv.Atoi(dbPort)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, mysqlPort, dbName)
	fmt.Println(dataSource)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(40 * time.Minute)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Storage) Get() *sql.DB {
	return s.db
}

func (s *Storage) Close() {
	s.db.Close()
}
