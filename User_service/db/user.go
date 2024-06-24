package db

import (
	"database/sql"
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
		"TYPE":      "User",
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
func (r *UserTypeRepo) CheckUser(email string) (string, error) {
	var pass string

	queryString, args, err := GetQueryBuilder().
		Select("password").
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
		return pass, err
	}
	err = GetReadDB().Get(&pass, queryString, args...)
	return pass, err
}
func (r *UserTypeRepo) Update(usr User) error {

	// Build the insert query
	query, args, err := GetQueryBuilder().
		Update(r.table).
		Set("name", usr.Name).
		Set("email", usr.Email).
		Set("type", usr.Type).
		Set("is_active", usr.Is_active).
		Where(sq.Eq{"id": usr.Id}).
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
	err = GetWriteDB().QueryRow(query, args...).Err()
	return err
}
func (r *UserTypeRepo) ReadUser(id int) (User, error) {

	// Build the select query
	query, args, err := GetQueryBuilder().
		Select("id", "name", "email", "type", "is_active").
		From(r.table).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("Failed to create user select query: %v, query: %s, args: %v\n", err, query, args)
		return User{}, err
	}

	var usr User
	err = GetWriteDB().QueryRow(query, args...).Scan(
		&usr.Id,
		&usr.Name,
		&usr.Email,
		&usr.Type,
		&usr.Is_active,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with id: %d\n", id)
			return User{}, nil // Return nil error for not found
		}
		log.Printf("Failed to execute query: %v\n", err)
		return User{}, err
	}

	return usr, nil
}
