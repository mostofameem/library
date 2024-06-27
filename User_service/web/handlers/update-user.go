package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"user_service/db"
	"user_service/logger"
	"user_service/web/utils"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	bodyparam, err := extractBodyParams(r)
	if err != nil {
		slog.Error("Failed to get body Params", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": bodyparam,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	// Check if "id" is present and not empty
	if id, ok := bodyparam["id"].(float64); !ok || id == 0 {
		slog.Error("no ID fild", logger.Extra(map[string]any{
			"error":   ok,
			"payload": bodyparam,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GetUserTypeRepo().UpdateUser(bodyparam)
	if err != nil {
		slog.Error("Update Failed", logger.Extra(map[string]any{
			"error":   err,
			"payload": bodyparam,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendData(w, "Update successful")

}

func extractBodyParams(r *http.Request) (map[string]any, error) {

	var bodydata map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodydata)
	if err != nil {
		return nil, err
	}
	return bodydata, nil

}
