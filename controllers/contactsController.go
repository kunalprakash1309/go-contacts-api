package controllers

import (
	"encoding/json"
	"github.com/go-contacts-api/models"
	_ "github.com/go-contacts-api/models"
	"github.com/go-contacts-api/utils"
	"net/http"
)

func CreateContact(w http.ResponseWriter, r *http.Request){

	// Grab the id of the user that send the request
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	utils.Respond(w, resp)
}

func GetContactsFor(w http.ResponseWriter, r *http.Request){

	id := r.Context().Value("user").(uint)

	data := models.GetContacts(id)
	resp := utils.Message(true, "Success")
	resp["data"] = data
	utils.Respond(w, resp)
}