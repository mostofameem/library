package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"user_service/db"
	"user_service/logger"
	"user_service/web/middlewares"
	"user_service/web/utils"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	userinfo, _ := db.GetUserTypeRepo().CheckUser(user.Email)

	accessToken, err := middlewares.GenerateToken(userinfo)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, err)
	}

	if userinfo.Pass == hashPassword(user.Password) {
		log.Println(accessToken)
		utils.SendBothData(w, userinfo.Type, accessToken)
	} else {
		utils.SendError(w, http.StatusUnauthorized, fmt.Errorf("worng Username/password"))
	}
}
