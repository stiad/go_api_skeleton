package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stiad/Api_Skeleton/src/app"
)

func main() {
	server := app.NewServer()
	server.LocalDev("8080")
	// server.Serve("8080")
}
