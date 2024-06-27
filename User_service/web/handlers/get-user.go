package handlers

import (
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"user_service/db"
	"user_service/logger"
	"user_service/web/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params, err := extractParams(r.URL.Query())
	if err != nil {
		slog.Error("Failed to extract query params", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": params,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	log.Println(params)

	UserList, err := db.GetUserTypeRepo().GetUser(params)
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
