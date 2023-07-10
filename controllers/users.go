package controllers

import (
	"basic_api/config"
	"basic_api/models"
	"basic_api/util"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "json/application")
	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)
	var user models.Users
	DB := config.Dbmigration()
	DB.Where("email=?", data["email"]).First(&user)
	if user.ID == 0 {
		json.NewEncoder(w).Encode("user not found")
		return
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		json.NewEncoder(w).Encode("incorrect password")
		return
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))

	if err != nil {
		panic("error")
	}
	/*	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	})*/
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode("sign in success")

}
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var data map[string]string

	json.NewDecoder(r.Body).Decode(&data)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.Users{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	db := config.Dbmigration()
	db.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	/*http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HttpOnly: true,
	})*/
	/*cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)*/
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode("logout sucess")
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	cookie, err := r.Cookie("jwt")
	if err != nil {
		panic(err.Error())
	}
	//tpken := cookie.Value
	id, err := util.ParseJwt(cookie.Value)
	if err != nil {
		json.NewEncoder(w).Encode("something went wrong")
		return
	}

	var user models.Users
	db := config.Dbmigration()
	db.Where("id=?", id).First(&user)
	json.NewEncoder(w).Encode(user)
}
func Deleteuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	cookie, err := r.Cookie("jwt")
	if err != nil {
		panic(err.Error())
	}
	//tpken := cookie.Value
	id, _ := util.ParseJwt(cookie.Value)
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)
	db := config.Dbmigration()
	db.Delete(&user, id)
	json.NewEncoder(w).Encode("user is deleted")
}
