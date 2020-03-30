package main

import (
	"flag"
	"github.com/landmaj/shrt/app"
	"log"
)

func main() {
	createDB := flag.Bool("db", false, "create database tables")
	flag.Parse()
	if *createDB {
		log.Println("Creating database tables...")
		app.CreateDatabase(app.NewDatabase())
	} else {
		app.Run()
	}
}
