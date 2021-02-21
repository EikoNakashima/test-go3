package main

import (
	"github.com/sample/sample-go-api/db"
	"github.com/sample/sample-go-api/server"
)

func main() {
	db.Init()
	server.Init()
}
