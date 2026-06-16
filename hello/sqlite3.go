package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)
type Organization map[string]interface{}

type Device struct {
	ID  string `db:"id"`
	URL string `db:"url"`
	Organization Organization `db:"organization"`
}

func setupDb() {
		db_name := os.Getenv("DB_NAME")
		db, err := sqlx.Open("sqlite3", db_name)
		if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

		createDevicesStmt := `
		drop table if exists devices;
		CREATE TABLE if not exists devices (
    id              TEXT PRIMARY KEY,
    url             TEXT NOT NULL
	);
		`
    //organization    TEXT NOT NULL,
	_, err = db.Exec(createDevicesStmt)

	if err != nil {
		log.Printf("%q: %s\n", err, createDevicesStmt)
		return
	}
}

func writeDevice(id string, url string) {
		db_name := os.Getenv("DB_NAME")
		db, err := sqlx.Open("sqlite3", db_name)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		res, err := db.Exec(`
			insert into devices (id, url) values (?, ?)
			`, id, url,
		)

		if err != nil {
			log.Printf("%q: %s\n", err, "insert into devices (id, url) values (?, ?)")
			return
		}

		fmt.Printf("device inserted result %v", res)

}

func readDevices() []Device{
		db_name := os.Getenv("DB_NAME")
		db, err := sqlx.Open("sqlite3", db_name)
		if err != nil {
			log.Fatal(err)
		}
		
		defer db.Close()
		var d []Device

		_ = db.Select(&d, "select * from devices")
		//fmt.Printf("read devices result %v", d)
		return d
}