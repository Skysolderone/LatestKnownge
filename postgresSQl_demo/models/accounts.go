package models

import (
	"os"
	"strings"

	"v1/utils"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (acc *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(acc.Email, "@") {
		return utils.Message(false, "email adderss is required"), false
	}
	if len(acc.Password) < 6 {
		return utils.Message(false, "password <6"), false
	}
	temp := &Account{}
	err := GetDb().Table("accounts").Where("email=?", acc.Email).
		First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}
	return utils.Message(false, "Requirement passed"), true
}

func (acc *Account) Create() map[string]interface{} {
	if resp, ok := acc.Validate(); !ok {
		return resp
	}
	hahspas, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	acc.Password = string(hahspas)
	GetDb().Create(acc)
	if acc.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}
	tk := &Token{UserId: acc.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Token = tokenString
	acc.Password = " "
	reponse := utils.Message(true, "Accoutn has been created")
	reponse["account"] = acc
	return reponse
}

func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDb().Table("accounts").Where("email=?", email).
		First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(false, "Connection error. Please retry")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { // Password does not match!
		return utils.Message(false, "Invalid login credentials. Please try again")
	}
	account.Password = " "
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	resp := utils.Message(true, "Loggined in")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {
	acc := &Account{}
	GetDb().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { // User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
