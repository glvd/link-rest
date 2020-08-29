package main

import (
	rest "github.com/glvd/link-rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goextension/log/zap"
)

func main() {
	zap.InitZapSugar()
	rest, err := rest.New(18080)
	if err != nil {
		panic(err)
	}
	rest.Start()
}
