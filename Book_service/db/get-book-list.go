package db

import (
	"book_service/web/utils"
	"log"

	sq "github.com/Masterminds/squirrel"
)

type BooksFilter struct {
	Title            string
	Author           string
	Genres           string
	Publication_date string
	Status           string
}

type BooksFilterParams struct {
	OperationFilters BooksFilter `json:"book_filters"`
	Page             int         `json:"page"`
	Limit            int         `json:"limit"`
	SortBy           string      `json:"sort_by"`
	SortOrder        string      `json:"sort_order"`
}

func (r *BookTypeRepo) GetBookList(params utils.PaginationParams) ([]Books, error) {
	var AllProducts []Books

	limit, offset := ConfigPageSize(params.Page, params.Limit)

	Query := GetQueryBuilder().
		Select("isbn", "title", "author", "genres", "quantity", "Publication_date", "Next_available", "Is_active").
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
		return []Books{}, err
	}

	log.Println(sql)

	err = GetDB().Select(&AllProducts, sql, args...)
	if err != nil {
		log.Println("Failed to get books:", err)
		return []Books{}, err
	}

	log.Println(AllProducts)
	return AllProducts, nil

}
func ConfigPageSize(page, limit int) (int, int) {
	PageLimit := 20
	Offset := 0

	PageLimit = min(limit, PageLimit)
	Offset = PageLimit * page

	return PageLimit, Offset
}
