package models

import (
	"github.com/go-contacts-api/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Contact struct{
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
}


func (contact *Contact) Validate() (map[string]interface{}, bool) {
	
	if contact.Name == ""{
		return utils.Message(false, "Contact name should be on the payload"), false
	}

	if contact.Phone == "" {
		return utils.Message(false, "Phone number should on the payload"), false
	}

	if contact.UserId <= 0 {
		return utils.Message(false, "User is not recognized"), false
	}

	// All the required parameters are present
	return utils.Message(true, "Success"), true
}

func (contact *Contact) Create() (map[string]interface{}) {
	
	if resp, ok := contact.Validate(); !ok{
		return resp
	}

	GetDB().Create(contact)

	resp := utils.Message(true, "Success")
	resp["contact"] = contact
	return resp
}