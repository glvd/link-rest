package scrape

import (
	"context"
	"github.com/glvd/link-rest/db"
	"github.com/goextension/log"
	"github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"gorm.io/gorm"
)

type Scraper interface {
}

type scrape struct {
	api *httpapi.HttpApi
	db  *gorm.DB
}

func New(path string) (Scraper, error) {
	api, err := httpapi.NewPathApi(path)
	if err != nil {
		return nil, err
	}
	dbcfg := db.ParseFromMap(nil)
	db, err := db.New(dbcfg)
	if err != nil {
		return nil, err
	}
	return &scrape{
		api: api,
		db:  db,
	}, nil
}

func (s *scrape) ParseHash(ctx context.Context, hash string) error {
	ls, err := s.api.Unixfs().Ls(ctx, path.New(hash))
	if err != nil {
		return err
	}

	for inf := range ls {
		log.Infof("%+v", inf)
	}
	return nil
}
