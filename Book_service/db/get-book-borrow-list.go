package db

import (
	"book_service/web/utils"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *BorrowTypeRepo) GetBookBorrowList(params utils.PaginationParams) ([]BorrowDetails, error) {
	var AllProducts []BorrowDetails

	limit, offset := ConfigPageSize(params.Page, params.Limit)

	Query := GetQueryBuilder().
		Select("book_title", "page_readed", "page_in_book", "issue_date", "return_date", "return_status").
		From(r.table)

	for k, v := range params.Filters {
		Query = Query.Where(sq.Eq{k: v})
	}
	// Apply limit and offset
	Query = Query.Limit(uint64(limit)).
		Offset(uint64(offset))

	// Execute the query
	sql, args, err := Query.ToSql()
	if err != nil {
		log.Println("Error building SQL:", err)
		return []BorrowDetails{}, err
	}

	err = GetDB().Select(&AllProducts, sql, args...)
	if err != nil {
		log.Println("Failed to get books:", err)
		return []BorrowDetails{}, err
	}
	return AllProducts, nil

}
