package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	v0 "github.com/glvd/link-rest/controller/v0"
	"github.com/glvd/link-rest/db"
	"github.com/glvd/link-rest/model"
	"gorm.io/gorm"
)

type Controller interface {
	Start() error
	Stop() error
}

type controller struct {
	ctx    context.Context
	cancel context.CancelFunc
	engine *gin.Engine
	cache  *persistence.InMemoryStore
	port   int
	serv   http.Server
	db     *gorm.DB
}

func (s *controller) Start() error {
	if err := model.Migration(s.db); err != nil {
		return err
	}
	s.registerHandle()
	s.serv.Handler = s.engine
	s.serv.Addr = fmt.Sprintf("0.0.0.0:%d", s.port)
	return s.serv.ListenAndServe()
}

func (s *controller) Stop() error {
	return s.serv.Close()
}

func (s *controller) registerHandle() {
	apiDocs(s.engine)
	groupV0 := s.engine.Group("/api/v0")
	v0.Register(s.db, groupV0, s.cache)
}

func New(port int) (Controller, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	dbcfg := db.ParseFromMap(nil)
	db, err := db.New(dbcfg)
	if err != nil {
		return nil, err
	}
	return &controller{
		ctx:    ctx,
		cancel: cancel,
		port:   port,
		serv:   http.Server{},
		engine: gin.Default(),
		cache:  persistence.NewInMemoryStore(time.Second),
		db:     db,
	}, nil
}
