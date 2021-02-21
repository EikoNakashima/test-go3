package main

import (
	"github.com/EikoNakashima/test-go3/db"
	"github.com/EikoNakashima/test-go3/server"
)

func main() {
	db.Init()
	server.Init()
}
