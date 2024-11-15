package controllers

import (
	"golang-backend/configs"
	"golang-backend/helpers"
	"golang-backend/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	
	
	passwordHash, err := helpers.HashPassword(password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Name:    name,
		Email:   email,
		Password: passwordHash,

	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "User has been registered", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var user models.User
	if err := configs.DB.First(&user, "email = ?", email).Error; err != nil {
		helpers.Response(w, 404, "Invalid Email or Password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, password); err != nil {
		helpers.Response(w, 404, "Invalid Email or Password", nil)
		return
	}

	token, err := helpers.CreateToken(&user)

	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Succesfully Login", token)
}