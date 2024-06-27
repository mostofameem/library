package db

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *UserTypeRepo) GetUser(conditions map[string]interface{}) ([]User, error) {

	// Build the select query
	queryBuilder := GetQueryBuilder().
		Select("id", "name", "email").
		From(r.table)

	for key, value := range conditions {
		queryBuilder = queryBuilder.Where(sq.Eq{key: value})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Printf("Failed to create user select query: %v, query: %s, args: %v\n", err, query, args)
		return []User{}, err
	}

	var usr []User
	err = GetReadDB().Select(&usr, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No users found: %v\n", err)
		} else {
			log.Printf("Failed to get users: %v\n", err)
		}
		return []User{}, err
	}

	return usr, nil

}
