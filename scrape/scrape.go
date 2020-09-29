package scrape

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/glvd/link-rest/db"
	"github.com/glvd/link-rest/model"
	"github.com/goextension/log"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

type Scraper interface {
	ParseHash(ctx context.Context, hash string) error
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

func NewScraper(api *httpapi.HttpApi, db *gorm.DB) Scraper {
	return &scrape{
		api: api,
		db:  db,
	}
}

func (s *scrape) ParseHash(ctx context.Context, hash string) error {
	ls, err := s.api.Unixfs().Ls(ctx, path.New(hash))
	if err != nil {
		return err
	}

	data := map[string]cid.Cid{}
	for inf := range ls {
		if (inf.Name == "info" || inf.Name == "media") && inf.Type == iface.TDirectory {
			data[inf.Name] = inf.Cid
			continue
		}
		return errors.New("found unknown file")
	}
	f := model.File{}
	m, b := data["media"]
	if !b {
		return errors.New("media not found")
	}
	f.RootHash = hash
	f.M3U8Path = strings.Join([]string{hash, "media"}, "/")
	f.M3U8Hash = m.String()
	f.M3U8Index = "media.m3u8"
	v, b := data["info"]
	if !b {
		return errors.New("info not found")
	}
	ls, err = s.api.Unixfs().Ls(ctx, path.IpfsPath(v))
	if err != nil {
		return err
	}

	for inf := range ls {
		if inf.Name == "nfo.json" {
			f.InfoHash = inf.Cid.String()
			f.InfoPath = strings.Join([]string{hash, "info", inf.Name}, "/")
		} else if inf.Name == "poster.jpg" {
			f.PosterHash = inf.Cid.String()
			f.PosterPath = strings.Join([]string{hash, "info", inf.Name}, "/")
		} else if inf.Name == "thumb.jpg" {
			f.ThumbHash = inf.Cid.String()
			f.ThumbPath = strings.Join([]string{hash, "info", inf.Name}, "/")
		} else {

		}
	}

	node, err := s.api.Unixfs().Get(ctx, path.New(f.InfoHash))
	if err != nil {
		log.Error("http get error", err)
		return err
	}

	file, b := node.(files.File)
	if !b {
		return errors.New("info is not a file")
	}

	info := model.Info{}
	infoData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error("info read error", err)
		return err
	}
	err = json.Unmarshal(infoData, &info)
	if err != nil {
		return err
	}
	log.Infof("info:%+v", info)
	log.Infof("file:%+v", f)
	media := model.Media{
		Info: info,
		File: f,
	}
	var count int64
	retCount := s.db.Model(model.File{}).Where(model.File{RootHash: f.RootHash}).Count(&count)
	if retCount.Error != nil {
		return retCount.Error
	}
	if count > 0 {
		log.Infof("file exist:%+v skip", f.RootHash)
		return nil
	}

	ret := s.db.Create(&media)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
