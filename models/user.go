package models

import "time"

type User struct { // default table name `users`
	Id         int
	Username   string
	Age        int
	Email      string
	AddTime    time.Time
	UpdateTime time.Time
}

func (c User) TableName() string {
	return "user"
}
