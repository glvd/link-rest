package model

import (
	"fmt"
	"github.com/glvd/link-rest/db"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var testdb *gorm.DB

func init() {
	cfg := db.ParseFromMap(nil)
	fmt.Println(cfg.String())
	engine, err := db.New(cfg)
	if err != nil {
		panic(err)
	}
	testdb = engine
}

func TestInsertMedia(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	err := Migration(testdb)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10000; i++ {
		//media := generateTestMedia("")
		//_, err = testdb.Create(media.File)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//_, err = testdb.Insert(media.Info)
		//if err != nil {
		//	t.Fatal(err)
		//}
		db := testdb.Create("media")
		if db.Error != nil {
			t.Fatal(db.Error)
		}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				db: testdb,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Migration(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Migration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
