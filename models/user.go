package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID   int `gorm:"primaryKey; autoIncrement"`
	Name string
}

func NewUser() *User {
	return &User{}
}

func (u *User) BuildName(name string) *User {
	u.Name = name
	return u
}

func (u *User) BuildId(id int) *User {
	u.ID = id
	return u
}
