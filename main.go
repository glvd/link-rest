package main

import (
	c "github.com/glvd/link-rest/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goextension/log/zap"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:18080
// @BasePath /api/v0
func main() {
	zap.InitZapSugar()

	rest, err := c.New(18080)
	if err != nil {
		panic(err)
	}
	rest.Start()
}
