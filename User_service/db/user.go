package db

import (
	"log"
	"log/slog"
	"user_service/logger"

	sq "github.com/Masterminds/squirrel"
)

type User struct {
	Id        int    `db:"id"`
	Name      string `db:"name" `
	Email     string `db:"email" `
	Pass      string `db:"password"`
	Type      string `db:"type"`
	Is_active string `db:"is_active"`
}
type UserTypeRepo struct {
	table string
}

var userTypeRepo *UserTypeRepo

func initUserTypeRepo() {
	userTypeRepo = &UserTypeRepo{
		table: "users",
	}
}

func GetUserTypeRepo() *UserTypeRepo {
	return userTypeRepo
}

var UserID int

func (r *UserTypeRepo) Create(usr User) error {
	// Prepare columns and values for the insert query
	columns := map[string]interface{}{
		"name":      usr.Name,
		"Email":     usr.Email,
		"password":  usr.Pass,
		"type":      "User",
		"is_active": false,
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
		Suffix("RETURNING id").
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

	// Execute the query and get the returned order_id
	err = GetWriteDB().QueryRow(query, args...).Scan(&UserID)
	return err
}
func (r *UserTypeRepo) CheckUser(email string) (User, error) {
	var returnuserinfo User
	queryString, args, err := GetQueryBuilder().
		Select("id", "password", "type").
		From(r.table).
		Where(sq.Eq{"email": email}).
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
		return returnuserinfo, err
	}
	err = GetReadDB().Get(&returnuserinfo, queryString, args...)
	return returnuserinfo, err
}

type UserParam struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Is_active string `json:"is_active"`
}

// Query := GetQueryBuilder().
// 		Select("isbn", "title", "total_page", "author", "genres", "quantity", "Publication_date", "Next_available", "Is_active").
// 		From(r.table)

//	for k, v := range params.Filters {
//		Query = Query.Where(sq.Eq{k: v})
//	}
func (r *UserTypeRepo) ReadUser(qry UserParam) ([]User, error) {

	Query := GetQueryBuilder().
		Select("id", "name", "email").
		From(r.table)

	if qry.Id != 0 {
		Query = Query.Where(sq.Eq{"id": qry.Id})
	}
	if qry.Type != "" {
		Query = Query.Where(sq.Eq{"type": qry.Type})
	}
	if qry.Is_active != "" {
		Query = Query.Where(sq.Eq{"is_active": qry.Is_active})
	}

	var user []User
	sql, args, err := Query.ToSql()
	if err != nil {
		log.Println("Error building SQL:", err)
		return []User{}, err
	}
	err = GetReadDB().Select(&user, sql, args...)
	if err != nil {
		log.Println("Failed to get books:", err)
		return []User{}, err
	}
	return user, nil
}

func (r *UserTypeRepo) ApproveUserRequest(q User) error {
	query, args, err := GetQueryBuilder().Update(r.table).
		Set("is_active", true).
		Where(sq.Eq{"id": q.Id}).ToSql()

	log.Println("hi")
	log.Println(query)

	if err != nil {
		slog.Error("Failed to make query", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": query,
		}))
		return err
	}

	err = GetWriteDB().QueryRow(query, args...).Err()
	return err
}

func (r *UserTypeRepo) RejectUserRequest(q User) error {
	query, args, err := GetQueryBuilder().Delete(r.table).
		Where(sq.Eq{"id": q.Id}).ToSql()
	if err != nil {
		slog.Error("Failed to Delete data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": query,
		}))
		return err
	}

	err = GetWriteDB().QueryRow(query, args...).Err()
	return err
}
