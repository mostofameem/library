package db

import (
	"book_service/logger"
	"database/sql"
	"log"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

type BorrowTypeRepo struct {
	table string
}

var borrowTypeRepo *BorrowTypeRepo

func initBorrowTypeRepo() {
	borrowTypeRepo = &BorrowTypeRepo{
		table: "borrow",
	}
}

func GetBorrowTypeRepo() *BorrowTypeRepo {
	return borrowTypeRepo
}

type BorrowDetails struct {
	User_id       int    `db:"user_id" json:"user_id"`
	Book_title    string `db:"book_title" json:"book_title"`
	Page_readed   int    `db:"page_readed" json:"page_readed"`
	Page_in_book  int    `db:"page_in_book" json:"total_page"`
	Issue_date    string `db:"issue_date" json:"issue_date"`
	Return_date   string `db:"return_date" json:"return_date"`
	Return_status string `db:"return_status" json:"return_status"`
	Is_active     bool   `db:"is_active" json:"is_active"`
}

func (r *BorrowTypeRepo) AddBorrowRequest(q BorrowDetails) error {
	// Prepare columns and values for the insert query
	columns := map[string]interface{}{
		"user_id":       q.User_id,
		"book_title":    q.Book_title,
		"page_readed":   0,
		"page_in_book":  q.Page_in_book,
		"issue_date":    q.Issue_date,
		"return_date":   q.Return_date,
		"return_status": false,
		"is_active":     false,
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

func (r *BorrowTypeRepo) ApproveUserRequest(q BorrowDetails) error {
	query, args, err := GetQueryBuilder().Update(r.table).
		Set("is_active", true).
		Where(sq.Eq{"user_id": q.User_id}).
		Where(sq.Eq{"book_title": q.Book_title}).ToSql()
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": query,
		}))
		return err
	}
	err = GetDB().QueryRow(query, args...).Err()
	return err
}

func (r *BorrowTypeRepo) RejectUserRequest(q BorrowDetails) error {
	query, args, err := GetQueryBuilder().Delete(r.table).
		Where(sq.Eq{"user_id": q.User_id}).
		Where(sq.Eq{"book_title": q.Book_title}).ToSql()
	if err != nil {
		slog.Error("Failed to Delete data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": query,
		}))
		return err
	}

	log.Println(query)

	err = GetDB().QueryRow(query, args...).Err()
	return err
}
func (r *BorrowTypeRepo) GetBorrowList(conditions map[string]interface{}) ([]BorrowDetails, error) {

	// Build the select query
	queryBuilder := GetQueryBuilder().
		Select("user_id", "book_title").
		From(r.table)

	for key, value := range conditions {
		queryBuilder = queryBuilder.Where(sq.Eq{key: value})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Printf("Failed to create user select query: %v, query: %s, args: %v\n", err, query, args)
		return []BorrowDetails{}, err
	}

	var borrowlist []BorrowDetails
	err = GetDB().Select(&borrowlist, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No users found: %v\n", err)
		} else {
			log.Printf("Failed to get users: %v\n", err)
		}
		return []BorrowDetails{}, err
	}

	return borrowlist, nil

}
