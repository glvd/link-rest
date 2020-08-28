package linkrest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/glvd/link-rest/db"
	v0 "github.com/glvd/link-rest/v0"
	"github.com/xormsharp/xorm"
	"net/http"
)

type Service interface {
	Start() error
	Stop() error
}

type service struct {
	ctx    context.Context
	cancel context.CancelFunc
	engine *gin.Engine
	port   int
	serv   http.Server
	db     *xorm.Engine
}

func (s *service) Start() error {
	s.registerHandle()
	s.serv.Handler = s.engine
	s.serv.Addr = fmt.Sprintf("0.0.0.0:%d", s.port)
	return s.serv.ListenAndServe()
}

func (s *service) Stop() error {
	return s.serv.Close()
}

func (s *service) registerHandle() {
	groupV0 := s.engine.Group("/api/v0")
	v0.RegisterV0(groupV0)
}

func New(port int) (Service, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	dbcfg := db.ParseFromMap(nil)
	dbeng, err := db.New(dbcfg)
	if err != nil {
		return nil, err
	}
	return &service{
		ctx:    ctx,
		cancel: cancel,
		port:   port,
		serv:   http.Server{},
		engine: gin.Default(),
		db:     dbeng,
	}, nil
}
