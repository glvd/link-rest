package restapi

import (
	"fmt"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/glvd/link-rest/db"
	"github.com/glvd/link-rest/restapi/common/controller"
	cm "github.com/glvd/link-rest/restapi/common/model"
	"net/http"
	"time"

	v0 "github.com/glvd/link-rest/restapi/v0/controller"
	v1 "github.com/glvd/link-rest/restapi/v1/controller"
)

type service struct {
	port      int
	apiPrefix string
	serv      http.Server
	c         *controller.Controller
}

type Service = service

func New(port int) (*Service, error) {
	eng := gin.Default()
	dbcfg := db.ParseFromMap(nil)
	gormdb, err := db.New(dbcfg)
	if err != nil {
		return nil, err
	}
	c := &controller.Controller{
		Engine: eng,
		Cache:  persistence.NewInMemoryStore(time.Second),
		DB:     gormdb,
	}

	return &service{
		port:      port,
		apiPrefix: "api",
		serv: http.Server{
			Handler: c.Engine,
			Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		},
		c: c,
	}, nil
}

func (s *service) init() error {
	if err := cm.Migration(s.c.DB); err != nil {
		return err
	}
	_ = v0.RegisterHandle(s.apiPrefix, s.c)
	_ = v1.RegisterHandle(s.apiPrefix, s.c)
	return nil
}

func (s *service) Start() error {
	if err := s.init(); err != nil {
		return err
	}
	return s.serv.ListenAndServe()
}

func (s *service) Stop() error {
	return s.serv.Close()
}
