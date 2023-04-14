package main

import (
	"final/database"
	"final/router"
)


func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run()

}