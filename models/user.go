package models

import (
	"time"
)

type User struct {
	GormModel
	Username        string        `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Email           string        `gorm:"not null;uniqueIndex" json:"email"`
	Password        string        `gorm:"not null" json:"password" form:"password"`
	Age             int           `gorm:"not null" json:"age" form:"age"`
	ProfileImageUrl string        `json:"profile_image_url" form:"profile_image_url"`
	Photos          []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments        []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	SocialMedias    []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

type CreateUserRequest struct {
	Age      int    `json:"age" form:"age" binding:"required,min=9"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
	Username string `json:"username" form:"username" binding:"required"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Username string `json:"username" form:"username" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type CreateUserResponse struct {
	ID       uint   `json:"id"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateUserResponse struct {
	ID        uint       `json:"id"`
	Age       int        `json:"age"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserPhotoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCommentResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserSocialMediaResponse struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}