package controllers

import (
	"net/http"
	"github.com/go-contacts-api/utils"
	"github.com/go-contacts-api/models"
	"encoding/json"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}

	resp := account.Create()
	utils.Respond(w, resp)

}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	utils.Respond(w, resp)
}