package models

import "time"

type Register struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type User struct {
	Id           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	HashPassword string    `json:"-" db:"hashed_pass"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
