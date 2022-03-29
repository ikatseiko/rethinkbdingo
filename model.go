package main

type Client struct {
	Id     string `json:"id" rethinkdb:"id,omitempty"`
	Name   string `json:"name" rethinkdb:"name"`
	Active bool   `json:"active" rethinkdb:"active"`
	Auto   bool   `json:"auto" rethinkdb:"auto"`
}

type User struct {
	Id     string     `json:"id" rethinkdb:"id,omitempty"`
	Login  string     `json:"login" rethinkdb:"login"`
	Status UserStatus `json:"status" rethinkdb:"status"`
}

type UserStatus string

const (
	UserNew    UserStatus = "NEW"
	UserActive UserStatus = "ACTIVE"
	UserBan    UserStatus = "BAN"
)
