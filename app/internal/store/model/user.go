package model

import (
	"time"
)

func NewUser() *User {
	return &User{Model: &Model{}}
}

type User struct {
	ID            int64     `json:"id" gorm:"column:id"`
	UserName      string    `json:"userName" gorm:"column:user_name"`
	Password      string    `json:"password" gorm:"column:password"`
	Email         string    `json:"email" gorm:"column:email"`
	DisplayName   string    `json:"displayName" gorm:"column:display_name"`
	DeptName      string    `json:"deptName" gorm:"column:dept_name"`
	Origin        int8      `json:"origin" gorm:"column:origin"`
	LastLoginTime time.Time `json:"lastLoginTime"  gorm:"column:last_login_time"`
	*Model
}

func (m *User) TableName() string {
	return "user"
}
func (m *User) IDField() string {
	return "id"
}
func (m *User) UserNameField() string {
	return "user_name"
}
func (m *User) PasswordField() string {
	return "password"
}
func (m *User) EmailField() string {
	return "email"
}
func (m *User) DisplayNameField() string {
	return "display_name"
}
func (m *User) DeptNameField() string {
	return "dept_name"
}
func (m *User) OriginField() string {
	return "origin"
}
func (m *User) LastLoginTimeField() string {
	return "last_login_time"
}
