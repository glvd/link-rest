package scrape

import (
	"context"
	"github.com/glvd/link-rest/db"
	cm "github.com/glvd/link-rest/restapi/common/model"
	_ "github.com/glvd/link-rest/restapi/v0/model"
	_ "github.com/glvd/link-rest/restapi/v1/model"

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
	testdb = engine.Debug()
	api, err := httpapi.NewLocalApi()
	if err != nil {
		return
	}
	testapi = api
}

func TestScrapeParseHash(t *testing.T) {
	err := cm.Migration(testdb)
	if err != nil {
		t.Fatal(err)
	}
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
				hash: "QmanRrB2r4TsT76sjHEsizQiwsG1Pa6gEGk3AXKPWaXHLa",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmU1AoYRMJvnH8T3aw9nQT8kdUX1X2AM3RQSUfbst7tXqT",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmcVgEdKXmjqiht9iLbJzNyRuLJwJWrmct52SunkUfkPGD",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmWEyKaka5VKc21DsaTZsaY1vLaqqMLRrt1ii5ZgXTVL9M",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmYBprprxEvjqgS37ceG6jJzq2ji4WYnttcCvuM75iB4RF",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmcDz4ZM3k4MWj17P8CPiHBvTJUwaPwsTMUTsXTy5h5xjr",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmaXKETjqUNEcYxf2QfMEQNLmWQsAiRu6CgRWbRTvAMdmK",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmS9uLgrv79brSzzcUqhyRf6CUR3KCxxZD4pZzRGJsGaS6",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmRZfEjYRgYRtAKfsDefqs85sZrhV8VKvGmWiRHh1hoWmV",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "Qmf29JBwVZvBcSx167FzHPPYJCy8gG6xPcDWsJGzPSizxG",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmWFXcwjmDKb2WRNzJs33RTG6KYyyNjAmfKiyL1ZSbLSMK",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmSr2B5Tbnks775dorMAq9T47b6ebFxbnGRe55dMK74XoY",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:  context.TODO(),
				hash: "QmW3vvz7pUGwAzdLDDUAiAkNE6yTUxWn7Som6trMzPeMPj",
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
