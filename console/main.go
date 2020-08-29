package main

import (
	rest "github.com/glvd/link-rest"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	rest, err := rest.New(18080)
	if err != nil {
		panic(err)
	}
	rest.Start()
}
