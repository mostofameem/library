package db

import (
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *UserTypeRepo) UpdateUser(conditions map[string]interface{}) error {

	// Build the select query
	queryBuilder := GetQueryBuilder().Update(r.table)

	for key, value := range conditions {
		if key == "id" {
			continue
		}
		queryBuilder = queryBuilder.Set(key, value)

	}
	queryBuilder = queryBuilder.Where(sq.Eq{"id": conditions["id"]})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Printf("Failed to create Update query: %v, query: %s, args: %v\n", err, query, args)
		return err
	}

	err = GetWriteDB().QueryRow(query, args...).Err()
	return err

}
