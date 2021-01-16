package controllers

import (
	"net/http"
	"github.com/go-contacts-api/models"
	"github.com/go-contacts-api/utils"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)

func CreateContact(w http.ResponseWriter, r *http.Request){
	user := r.Context().Value("userid")
}