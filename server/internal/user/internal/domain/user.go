package domain

import "time"

type User struct {
	ID        string
	Username  string
	Password  string
	RoleId    int
	Avatar    string
	Email     string
	Sex       string
	Company   string
	Hobby     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Page struct {
	PageNo   int64
	PageSize int64
}
