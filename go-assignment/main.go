package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "your_postgres_username"
	dbPassword = "your_postgres_password"
	dbName     = "your_postgres_dbname"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Database interface {
	GetBooks() ([]Book, error)
	SaveBook(book Book) error
}

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB() (Database, error) {
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	// connStr := "user=krunal password=krunalhingu dbname=citadel sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		);
	`

	_, err = db.Exec(createTable)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

func (db *PostgresDB) GetBooks() ([]Book, error) {
	rows, err := db.db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (db *PostgresDB) SaveBook(book Book) error {
	_, err := db.db.Exec("INSERT INTO books (title, author) VALUES ($1, $2)", book.Title, book.Author)
	if err != nil {
		return err
	}
	return nil
}

func setupRouter(db Database) *gin.Engine {
	r := gin.Default()

	r.GET("/books", func(c *gin.Context) {
		books, err := db.GetBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, books)
	})

	r.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if err := db.SaveBook(book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.Status(http.StatusCreated)
	})

	return r
}

func main() {
	db, err := NewPostgresDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	r := setupRouter(db)
	r.Run(":8080")
}
