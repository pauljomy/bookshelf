package data

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Cover    string `json:"cover"`    // computed from isbn13
	Progress int    `json:"progress"` // computed from shelf
}

type BookModel struct {
	db *pgxpool.Pool
}

func (m *BookModel) GetAll() ([]Book, error) {

	query := `select id, title, author, coalesce(isbn13, '') from books order by id`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := m.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		var isbn13 string

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &isbn13)
		if err != nil {
			return nil, err
		}

		if isbn13 != "" {
			book.Cover = fmt.Sprintf("https://covers.openlibrary.org/b/isbn/%s-M.jpg", isbn13)
		}

		books = append(books, book)
	}
	return books, nil
}
