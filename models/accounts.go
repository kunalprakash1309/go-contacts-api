package models

import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-contacts-api/utils"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Token struct{
	UserId uint
	jwt.StandardClaims
}

// Account to rep user account
type Account struct{
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`        // confuse
}

//Validate incoming user details
func (account *Account) Validate() (map[string]interface{}, bool){

	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}

	// Email must be unique
	temp := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}

	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}

	return utils.Message(true, "Requirement passed"), true

}

func (account *Account) Create() (map[string]interface{}) {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	// confuse
	GetDB().Create(account)
	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account")
	}

	//Create new JWT token for newly registred account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password"))) 
	account.Token = tokenString

	account.Password = "" // delete password

	response := utils.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) (map[string]interface{}) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return utils.Message(false, "Email address not found")
		}

		return utils.Message(false, "Connection error, Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword{
		return utils.Message(false, "Invalid Login Credentials")
	}

	// Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := utils.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u int) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	
	// User not found
	if acc.Email == ""{
		return nil
	}

	acc.Password = ""
	return acc
}
