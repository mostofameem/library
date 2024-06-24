package handlers

import (
	"book_service/db"
	"book_service/logger"
	"book_service/web/utils"
	"log/slog"
	"net/http"
)

const (
	defaultSortBy    = "created_at"
	defaultSortOrder = "desc"
)

type BooksFilter struct {
	Title          string `json:"title"`
	Author         string `json:"author"`
	Genres         string `json:"genres"`
	Next_available string `json:"status"`
}

type BooksFilterParams struct {
	OperationFilters BooksFilter `json:"book_filters"`
	Page             int         `json:"page"`
	Limit            int         `json:"limit"`
	SortBy           string      `json:"sort_by"`
	SortOrder        string      `json:"sort_order"`
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	//var params BooksFilterParams
	paginationParams := utils.GetPaginationParams(r, defaultSortBy, defaultSortOrder)
	//err := parseOperationFilters(&params)
	//log.Println(err)
	booklist, err := db.GetBookTypeRepo().GetBookList(paginationParams)
	if err != nil {
		slog.Error("Failed to GetBookList ", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": booklist,
		}))
	}
	utils.SendData(w, booklist)
}

// func parseOperationFilters(params *BooksFilterParams) error {
// 	return nil
// }
