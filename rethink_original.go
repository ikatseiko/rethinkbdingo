package main

import (
	"fmt"
	"log"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

const (
	tableUser   string = "Users"
	tableClient string = "clients"
	tablePE     string = "BAN"
)

func conect() *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "prepc",
		Username: "prepc",
		Password: "prepc",
	})
	if err != nil {
		return nil
	}
	return session
}

func AllClients(ss *r.Session) *[]Client {
	var clients []Client

	res, err := r.Table(tableClient).Run(ss)
	if err != nil {
		return &clients
	}
	res.All(&clients)

	return &clients
}

func AddClient(ss *r.Session, c *Client) bool {
	_, err := r.Table(tableClient).Insert(c).RunWrite(ss)
	return err == nil
}

func UpdateClient(ss *r.Session, c *Client) bool {
	_, err := r.Table(tableUser).Get(c.Id).Update(c).RunWrite(ss)
	return err == nil
}

func DelClient(ss *r.Session, id string) bool {
	err := r.Table(tableClient).Get(id).Delete().Exec(ss)
	return err == nil
}

func AllUsers(ss *r.Session) *[]User {
	var users []User

	res, err := r.Table(tableUser).Run(ss)
	if err != nil {
		return &users
	}
	res.All(&users)

	return &users
}

func AddUser(ss *r.Session, u *User) bool {
	_, err := r.Table(tableUser).Insert(u).RunWrite(ss)
	return err == nil
}

func UpdateUser(ss *r.Session, u *User) bool {
	_, err := r.Table(tableUser).Get(u.Id).Update(u).RunWrite(ss)
	return err == nil
}

func DelUser(ss *r.Session, id string) bool {
	err := r.Table(tableUser).Get(id).Delete().Exec(ss)
	return err == nil
}

func close(ses *r.Session) error {
	return ses.Close()
}

func RO_Example() {
	session, err := r.Connect(r.ConnectOpts{
		Address: "127.0.0.1:28015", // endpoint without http  http://localhost:8080/#dataexplorer
	})
	if err != nil {
		log.Fatalln(err)
	}

	res, err := r.Expr("Hello World").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)

	// Output:
	// Hello World
}
