package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"user_service/db"
	"user_service/logger"
	"user_service/web/utils"
)

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}
type Code struct {
	Code string `json:"code" validate:"required"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var newuser NewUser
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newuser,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	_, err = db.GetUserTypeRepo().CheckUser(newuser.Email)
	if err == nil {
		utils.SendData(w, "User Already Exists!! ")
		return
	}

	var usr db.User
	usr.Name = newuser.Name
	usr.Email = newuser.Email
	usr.Pass = hashPassword(newuser.Password)

	err = utils.Validate(usr)
	if err != nil {
		slog.Error("Failed to validate user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": usr,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	err = db.GetUserTypeRepo().Create(usr)
	if err != nil {
		log.Println(err)
		slog.Error("Failed to insert user db ", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": usr,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	utils.SendData(w, "Wait for Admin confirmation")
	log.Println(db.UserID)

}
func UrlOperation(r string) (db.UserParam, error) {

	var item db.UserParam
	parsedUrl, err := url.Parse(r)
	if err != nil {
		return item, err
	}
	queryParams := parsedUrl.Query()
	item.Id, err = strconv.Atoi(queryParams.Get("id"))
	if err != nil {
		return db.UserParam{}, fmt.Errorf("string to int conversion failed")
	}
	item.Type = queryParams.Get("type")
	item.Is_active = queryParams.Get("is_active")

	err = utils.Validate(item)
	if err != nil {
		return db.UserParam{}, err
	}
	return item, nil
}
func hashPassword(pass string) string {

	h := sha1.New()
	h.Write([]byte(pass))
	hashValue := h.Sum(nil)
	return hex.EncodeToString(hashValue)
}
func ApproveUserRequest(w http.ResponseWriter, r *http.Request) {
	var info db.User
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		slog.Error("Failed to get user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": info,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GetUserTypeRepo().ApproveUserRequest(info)
	if err != nil {
		slog.Error("Failed to update data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": info,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendData(w, "Request Accepted")
}
func RejectUserRequest(w http.ResponseWriter, r *http.Request) {
	var info db.User
	params, err := extractParams(r.URL.Query())
	if err != nil {
		slog.Error("Failed to extract query params", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": params,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	info.Id = params["id"].(int)

	err = db.GetUserTypeRepo().RejectUserRequest(info)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendData(w, "Success")
}
