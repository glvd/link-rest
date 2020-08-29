package model

import (
	"testing"

	"github.com/glvd/link-rest/db"
	"github.com/goextension/tool"
	"github.com/google/uuid"
	"github.com/xormsharp/xorm"
)

var testdb *xorm.Engine

func init() {
	engine, _ := db.New(nil)
	testdb = engine
}

func generateTestMedia(id string) *Media {
	if id == "" {
		id = uuid.New().String()
	}
	return &Media{
		BaseModel: BaseModel{
			ID: id,
		},
		VideoNo:      "no_" + tool.GenerateRandomString(6, tool.RandomNum),
		Intro:        "intro_" + tool.GenerateRandomString(32),
		Alias:        nil,
		ThumbHash:    "hash_" + tool.GenerateRandomString(16),
		PosterHash:   "hash_" + tool.GenerateRandomString(16),
		SourceHash:   "hash_" + tool.GenerateRandomString(16),
		M3U8Hash:     "hash_" + tool.GenerateRandomString(16),
		Key:          "",
		M3U8:         "",
		Role:         nil,
		Director:     "",
		Systematics:  "",
		Season:       "",
		TotalEpisode: "",
		Episode:      "",
		Producer:     "",
		Publisher:    "",
		Type:         "",
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
	}
}

func TestInsertMedia(t *testing.T) {
	type args struct {
		db *xorm.Engine
	}
	for i := 0; i < 100; i++ {
		media := generateTestMedia("")
		_, err := testdb.Insert(media)
		if err != nil {
			t.Fatal(err)
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
