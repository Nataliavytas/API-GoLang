package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", ":memory:")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error)
	}

	createSchema(db)

	//routing things
	r := gin.Default()
	r.GET("/users/:name", getUsersHandler)
	r.POST("/users", addUserHandler)
	r.Run()
}

func createSchema(db *sqlx.DB) {
	schema := `CREATE TABLE user (
		id integer primary autoincremental,
		name varchar(56));`
	db.Exec(schema)
}

func getUsersHandler(c *gin.Context) {
	name := c.Param("name")
	lastname := c.Query("lastname")
	c.JSON(200, gin.H{
		"name": name + " " + lastname,
	})
}

func addUserHandler(c *gin.Context) {
	c.JSON(201, gin.H{
		"name":     "Joe",
		"lastname": "Doe",
	})
}
