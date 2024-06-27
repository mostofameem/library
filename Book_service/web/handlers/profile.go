package handlers

import (
	"book_service/db"
	"book_service/logger"
	"book_service/web/middlewares"
	"book_service/web/utils"
	"fmt"
	"log/slog"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	user_id, err := middlewares.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		slog.Error("Failed to extract Id from Header", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user_id,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	// Append user_id to the URL query parameters
	query := r.URL.Query()
	query.Add("user_id", fmt.Sprintf("%v", user_id))
	r.URL.RawQuery = query.Encode()

	//var params BooksFilterParams
	paginationParams := utils.GetPaginationParams(r, defaultSortBy, defaultSortOrder)
	//err := parseOperationFilters(&params)
	//log.Println(err)
	BookList, err := db.GetBorrowTypeRepo().GetBookBorrowList(paginationParams)
	if err != nil {
		slog.Error("Failed to GetBookList ", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": BookList,
		}))
	}
	utils.SendData(w, BookList)
}
