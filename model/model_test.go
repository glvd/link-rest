package model

import (
	"github.com/google/uuid"
	"github.com/xormsharp/xorm"
)

type Repository interface {
	Get(id uuid.UUID) (*Media, error)
	Create(id uuid.UUID, name string) error
}

type repo struct {
	db *xorm.Engine
	*Media
}

func (p *repo) Create(id uuid.UUID, vn string) error {
	media := &Media{
		BaseModel: BaseModel{
			ID: id.String(),
		},
		VideoNo: vn,
	}
	if _, err := p.db.Insert(media); err != nil {
		return err
	}
	return nil
}

func (p *repo) Get(id uuid.UUID) (*Media, error) {
	media := new(Media)

	err := p.db.Where("id = ?", id.String()).Find(media)

	return media, err
}
