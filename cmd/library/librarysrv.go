package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Nataliavytas/API-GoLang/internal/config"
	"github.com/Nataliavytas/API-GoLang/internal/database"
	"github.com/Nataliavytas/API-GoLang/internal/service/library"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := library.New(db, cfg)
	httpService := library.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS library (
			id integer primary key autoincrement,
			title varchar,
			author varchar, 
			price integer);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertBook := `INSERT INTO library (title, author, price) VALUES (?, ?, ?)`
	title := "Heartstopper"
	author := "Alice Olsman"
	price := 700
	db.MustExec(insertBook, title, author, price)
	return nil
}
