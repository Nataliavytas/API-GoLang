package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Nataliavytas/API-GoLang/internal/config"
	"github.com/Nataliavytas/API-GoLang/internal/database"
	"github.com/Nataliavytas/API-GoLang/internal/service/chat"
	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := chat.New(db, cfg)

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
	schema := `CREATE TABLE IF NOT EXISTS messages (
			id integer primary key autoincrement,
			text varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Message numer %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}
