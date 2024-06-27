package handlers

import (
	"book_service/db"
	"book_service/logger"
	"book_service/web/middlewares"
	"book_service/web/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	var BookOrder db.BorrowDetails
	err := json.NewDecoder(r.Body).Decode(&BookOrder)
	if err != nil {
		slog.Error("Failed to get Body data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": BookOrder,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	BookOrder.User_id, err = middlewares.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": BookOrder.User_id,
		}))
		utils.SendError(w, http.StatusInternalServerError, fmt.Errorf("error Parsing id from jwt token"))
		return
	}

	// Get the current time
	now := time.Now()
	BookOrder.Issue_date = now.Format("2006-01-02") // DD-MM-YYYY format

	err = db.GetBorrowTypeRepo().AddBorrowRequest(BookOrder)
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": BookOrder,
		}))
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}
	utils.SendData(w, "Request Added Successful")
}

func ApproveBorrowRequest(w http.ResponseWriter, r *http.Request) {
	var info db.BorrowDetails
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": info,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GetBorrowTypeRepo().ApproveUserRequest(info)
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": info,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendData(w, "Request Accepted")
}
func RejectBorrowRequest(w http.ResponseWriter, r *http.Request) {
	var info db.BorrowDetails
	params, err := extractParams(r.URL.Query())
	if err != nil {
		slog.Error("Failed to extract query params", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": params,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	info.User_id = params["user_id"].(int)
	info.Book_title = params["book_title"].(string)

	err = db.GetBorrowTypeRepo().RejectUserRequest(info)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendData(w, "Request Rejected")
}
func GetBorrowRequest(w http.ResponseWriter, r *http.Request) {
	params, err := extractParams(r.URL.Query())
	if err != nil {
		slog.Error("Failed to extract query params", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": params,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	UserList, err := db.GetBorrowTypeRepo().GetBorrowList(params)
	if err != nil {
		slog.Error("Failed to get Pending user", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": UserList,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendData(w, UserList)
}

// extractParams extracts query parameters from a URL and returns a map of conditions
func extractParams(values url.Values) (map[string]interface{}, error) {
	conditions := make(map[string]interface{})

	for key, values := range values {
		if len(values) > 0 {
			// Convert string values to appropriate types if necessary
			switch key {
			case "id":
				if id, err := strconv.Atoi(values[0]); err == nil {
					conditions[key] = id
				}
			case "user_id":
				if id, err := strconv.Atoi(values[0]); err == nil {
					conditions[key] = id
				}
			case "is_active":
				if isActive, err := strconv.ParseBool(values[0]); err == nil {
					conditions[key] = isActive
				}
			default:
				conditions[key] = values[0]
			}
		}
	}

	return conditions, nil
}
