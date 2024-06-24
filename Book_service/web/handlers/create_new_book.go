package handlers

import (
	"book_service/db"
	"book_service/logger"
	"book_service/web/utils"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func Addbook(w http.ResponseWriter, r *http.Request) {
	var newbook db.Books
	err := json.NewDecoder(r.Body).Decode(&newbook)
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newbook,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.Validate(newbook)
	if err != nil {
		slog.Error("Failed to validate user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newbook,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GetBookTypeRepo().CheckBookExists(newbook.Isbn)
	if err == nil {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("book Already Exists"))
		return
	}

	err = db.GetBookTypeRepo().Create(newbook)
	if err != nil {
		log.Println(err)
		slog.Error("Failed to insert user db ", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newbook,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	utils.SendData(w, "Book Added Successful")

}
