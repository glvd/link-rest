package scrape

import (
	"context"
	"github.com/glvd/link-rest/db"
	"github.com/goextension/log/zap"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"gorm.io/gorm"
	"testing"
)

var testdb *gorm.DB
var testapi *httpapi.HttpApi

func init() {
	zap.InitZapSugar()
	cfg := db.ParseFromMap(nil)
	engine, err := db.New(cfg)
	if err != nil {
		panic(err)
	}
	testdb = engine

	api, err := httpapi.NewLocalApi()
	if err != nil {
		return
	}
	testapi = api
}

func Test_scrape_ParseHash(t *testing.T) {
	type fields struct {
		api *httpapi.HttpApi
		db  *gorm.DB
	}
	type args struct {
		ctx  context.Context
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmUfgJNGaBHEXJ8X5hv9VntZXRsmunKPJn9uXyDULmKg1P",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scrape{
				api: testapi,
				db:  testdb,
			}
			if err := s.ParseHash(tt.args.ctx, tt.args.hash); (err != nil) != tt.wantErr {
				t.Errorf("ParseHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
