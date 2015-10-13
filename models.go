package main

import "gopkg.in/gorp.v1"

type Base struct {
	Db *gorp.DbMap
}

type UserJSON struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type User struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	RegistrationDate int64  `json:"-" db:"registration_date"`
}
