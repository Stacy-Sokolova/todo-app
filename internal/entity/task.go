package entity

import "time"

type Task struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	Created_at  time.Time `json:"created_at" db:"created_at"`
	Updated_at  time.Time `json:"updated_at" db:"updated_at"`
}

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InsertInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type AuthInput struct {
	Username string
	Password string
}
