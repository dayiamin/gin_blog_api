package models

import (
	
)

type User struct {
	BaseModel
	UserName  string `gorm:"unique" json:"user_name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-"`
	UserProfile  UserProfile  `gorm:"constraint:OnDelete:CASCADE;"`
	Posts     []Post         `gorm:"constraint:OnDelete:CASCADE;"`
	PostComment  []PostComment
}

type RegisterUsers struct {
	UserName string `json:"user_name" binding:"min=3,max=100,required"`
	Email    string `json:"email" binding:"min=3,max=100,required,email"`
	Password string `json:"password" binding:"min=6,max=100,required"`
}


type UserProfile struct {
	BaseModel
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	Bio			string	`json:"bio"`
	ProfilePic	string	`json:"profile_pic"`
	UserID uint		
}

type UserProfileRegister struct {
	FirstName	string	`json:"first_name" binding:"max=100"`
	LastName	string	`json:"last_name" binding:"max=100"`
	Bio			string	`json:"bio" binding:"max=400"`
	ProfilePic	string	`json:"profile_pic"`
}

type UserLoginRequest struct {
	Credential    string `json:"credential" binding:"required"`
	Password string `json:"password" binding:"required"`
}