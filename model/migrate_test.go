package model

import (
	"github.com/glvd/link-rest/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goextension/tool"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"testing"
)

var testdb *gorm.DB

func init() {
	cfg := db.ParseFromMap(nil)
	engine, err := db.New(cfg)
	if err != nil {
		panic(err)
	}
	testdb = engine
}

func generateTestMedia(id string) *Media {
	if id == "" {
		id = uuid.New().String()
	}

	root := "hash_" + tool.GenerateRandomString(32)
	return &Media{
		BaseModel: BaseModel{
			ID: id,
		},

		Info: Info{
			VideoNo:      "no_" + tool.GenerateRandomString(6, tool.RandomNum),
			Intro:        "intro_" + tool.GenerateRandomString(32),
			Alias:        nil,
			Role:         nil,
			Director:     "",
			Systematics:  "",
			Season:       "",
			TotalEpisode: "",
			Episode:      "",
			Producer:     "",
			Publisher:    "",
			MediaType:    "",
			Format:       "",
			Language:     "",
			Caption:      "",
			Group:        "",
			Index:        "",
			ReleaseDate:  "",
			Sharpness:    "",
			Series:       "",
			Tags:         nil,
			Length:       "",
			Sample:       nil,
			Uncensored:   false,
		},
		File: File{
			RootHash:   root,
			KeyPath:    "",
			KeyHash:    "",
			ThumbPath:  root + "/info/thumb.jpg",
			ThumbHash:  "hash_" + tool.GenerateRandomString(32),
			PosterPath: root + "/info/poster.jpg",
			PosterHash: "hash_" + tool.GenerateRandomString(32),
			SourcePath: "",
			SourceHash: "hash_" + tool.GenerateRandomString(32),
			M3U8Index:  "media.m3u8",
			M3U8Path:   root + "/media",
			M3U8Hash:   "hash_" + tool.GenerateRandomString(32),
		},
	}

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
		media := generateTestMedia("")
		//_, err = testdb.Create(media.File)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//_, err = testdb.Insert(media.Info)
		//if err != nil {
		//	t.Fatal(err)
		//}
		db := testdb.Create(media)
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
