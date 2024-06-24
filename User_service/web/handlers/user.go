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
type UserParam struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Is_active string `json:"is_active"`
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
func Approved(w http.ResponseWriter, r *http.Request) {

	params, err := UrlOperation(r.URL.String())
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	user, err := db.GetUserTypeRepo().ReadUser(params.Id)
	if err != nil {
		slog.Error("Failed to Read User ", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	if params.Type != "" {
		user.Type = params.Type
	}
	if params.Is_active != "" {
		user.Is_active = params.Is_active
	}
	err = db.GetUserTypeRepo().Update(user)
	if err != nil {
		log.Println(err)
		slog.Error("Failed to Update User", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendData(w, "Update User Successful")

}
func UrlOperation(r string) (UserParam, error) {

	var item UserParam
	parsedUrl, err := url.Parse(r)
	if err != nil {
		return item, err
	}
	queryParams := parsedUrl.Query()
	item.Id, err = strconv.Atoi(queryParams.Get("id"))
	if err != nil {
		return UserParam{}, fmt.Errorf("string to int conversion failed")
	}
	item.Type = queryParams.Get("type")
	item.Is_active = queryParams.Get("is_active")

	err = utils.Validate(item)
	if err != nil {
		return UserParam{}, err
	}
	return item, nil
}
func hashPassword(pass string) string {

	h := sha1.New()
	h.Write([]byte(pass))
	hashValue := h.Sum(nil)
	return hex.EncodeToString(hashValue)
}
