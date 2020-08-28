package model

import (
	"github.com/xormsharp/xorm"
	"math"
	"net/url"
	"strconv"
)

var DefaultPaginatorLimit = 10

type Paginator struct {
	Current   int
	Limit     int
	Data      interface{}
	Total     int64
	TotalPage int
}

type Counter interface {
	Count(session *xorm.Session) (int64, error)
}

func NewPage(v interface{}) *Paginator {
	return &Paginator{
		Current:   0,
		Limit:     0,
		Data:      v,
		Total:     0,
		TotalPage: 0,
	}
}

func (p *Paginator) Parse(vals url.Values) *Paginator {
	var err error
	limit := vals.Get("limit")
	p.Limit, err = strconv.Atoi(limit)
	if err != nil {
		p.Limit = DefaultPaginatorLimit
	}
	current := vals.Get("current")
	p.Current, err = strconv.Atoi(current)
	if err != nil {
		p.Current = 0
	}
	return p
}

func (p *Paginator) Find(session *xorm.Session) (*Paginator, error) {
	count, err := session.Count()
	if err != nil {
		return nil, err
	}
	p.Total = count
	p.TotalPage = 0
	if p.Total != 0 {
		p.TotalPage = int(math.Ceil(float64(p.Total) / float64(p.Limit)))
	}

	if p.Current >= p.TotalPage {
		p.Current = 0
	}
	err = session.Limit(p.Limit, p.Current*p.Limit).Find(p.Data)
	if err != nil {
		return nil, err
	}
	return p, nil
}
