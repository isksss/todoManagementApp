package main

// Header
type Content struct {
	userData Users
}

// user data
type Users struct {
	// uuid int
	name string
	email string
	password string
}

// database
type DbLogin struct {
	id       string
	password string
	host     string
	port     string
	dbname   string
}
