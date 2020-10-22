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

var pinData = []string{
	"bafybeif7b4txswioozvfwfcusyfo4qcojxoavzjsofzvxbbw5i2xvdhp54",
	"bafybeiglokwfpgral7ac6kogpowsbhzb4bawuvllb7s5v4f5vxsaierydy",
	"bafybeihyzaqplcngmartscvimfncjmcydmkrycoc77rzs7djeaslo7p43e",
	"bafybeidtbdcc3shjgzjxar7cqvqqs6eifyduv4ohklo6k7d4izcsq6kkda",
	"bafybeich4yzfoh7h2z2ra5se5datwhxdhiwokxvymzb6dpsld5jpbjinb4",
	"bafybeiajveygfkkto6676xiq4dtirxkgmylzytqzrqmgro64ge2scrcheu",
	"bafybeiciuhnkkuh6fht7thtku6esi3wuue56qwm3uuqanuowbwyylbtwoe",
	"bafybeih526hqhpy7hid3jfu645qc75ywlvtyobbry7vhsbevmhtcxncxfe",
	"bafybeiansxz7ci5xn5xfr7ecmj2hbbrn77et45gcnw2rmtrr4pybz2vwxq",
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
		ctx context.Context
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
				ctx: context.TODO(),
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
			for i := range pinData {
				if err := s.ParseHash(tt.args.ctx, pinData[i]); (err != nil) != tt.wantErr {
					t.Errorf("ParseHash() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
