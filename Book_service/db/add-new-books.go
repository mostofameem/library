package db

import (
	"book_service/logger"
	"log"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

type BookTypeRepo struct {
	table string
}

var bookTypeRepo *BookTypeRepo

func initBookTypeRepo() {
	bookTypeRepo = &BookTypeRepo{
		table: "books",
	}
}

func GetBookTypeRepo() *BookTypeRepo {
	return bookTypeRepo
}

type Books struct {
	Isbn             int    `json:"isbn" validate:"required"`
	Title            string `json:"title" validate:"required"`
	Author           string `json:"author" validate:"required"`
	Genres           string `json:"genres" validate:"required"`
	Quantity         int    `json:"quantity" validate:"required"`
	Publication_date string `json:"publication_date" validate:"required"`
	Next_available   string `json:"next_available"`
	Is_active        string `json:"is_active"`
}

func (r *BookTypeRepo) Create(book Books) error {
	log.Println(book)

	// Prepare columns and values for the insert query
	columns := map[string]interface{}{
		"isbn":             book.Isbn,
		"title":            book.Title,
		"author":           book.Author,
		"genres":           book.Genres,
		"quantity":         book.Quantity,
		"Publication_date": book.Publication_date,
		"Next_available":   "available",
		"Is_active":        "TRUE",
	}
	var colNames []string
	var colValues []any

	for colName, colVal := range columns {
		colNames = append(colNames, colName)
		colValues = append(colValues, colVal)
	}

	// Build the insert query
	query, args, err := GetQueryBuilder().
		Insert(r.table).
		Columns(colNames...).
		Values(colValues...).
		ToSql()
	if err != nil {
		slog.Error(
			"Failed to create New user insert query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": query,
				"args":  args,
			}),
		)
		return err
	}

	err = GetDB().QueryRow(query, args...).Err()
	return err
}

func (r *BookTypeRepo) CheckBookExists(isbn int) error {
	title := ""

	queryString, args, err := GetQueryBuilder().
		Select("title").
		From(r.table).
		Where(sq.Eq{"isbn": isbn}).
		ToSql()
	if err != nil {
		slog.Error(
			"Failed to create query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": queryString,
				"args":  args,
			}),
		)
		return err
	}

	err = GetDB().Get(&title, queryString, args...)
	return err
}
