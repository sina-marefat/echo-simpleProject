package main

import (
	"myProject/db"
	"myProject/server/http"
)

func main() {
	db.Init()
	http.Serve()
}
