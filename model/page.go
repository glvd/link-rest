package model

import (
	"github.com/goextension/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

var DefaultPaginatorPerPage = 10

type Paginator struct {
	CurrentPage  int
	LastPage     int
	PerPage      int
	Data         interface{}
	Total        int64
	FirstPageURL string
	LastPageURL  string
	NextPageURL  string
	PrevPageURL  string
	Path         string
	From         int
	To           int
}

func defaultPage(v interface{}) *Paginator {
	return &Paginator{
		CurrentPage: 1,
		Data:        v,
	}
}

func Page(r *http.Request, v interface{}) *Paginator {
	p := defaultPage(v)
	return p.parse(r)
}

func (p *Paginator) parse(r *http.Request) *Paginator {
	var err error
	urls := r.URL.Query()
	perPage := urls.Get("per_page")
	p.PerPage, err = strconv.Atoi(perPage)
	if err != nil {
		p.PerPage = DefaultPaginatorPerPage
	}
	current := urls.Get("page")
	p.CurrentPage, err = strconv.Atoi(current)
	if err != nil {
		p.CurrentPage = 1
	}

	p.Path = r.Host + r.URL.Path
	return p
}

func (p *Paginator) Find(db *gorm.DB) (*Paginator, error) {
	var count int64
	tx := db.Count(&count)
	log.Infow("count", "error", tx.Error, "count", count)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if p.CurrentPage <= 0 || p.CurrentPage > p.LastPage {
		p.CurrentPage = 1
	}

	if count == 0 {
		return p, nil
	}

	p.Total = count
	p.From = (p.CurrentPage - 1) * p.PerPage
	p.To = p.From + p.PerPage
	p.LastPage = int(math.Ceil(float64(p.Total) / float64(p.PerPage)))
	p.NextPageURL = p.next()
	p.PrevPageURL = p.prev()
	p.LastPageURL = p.last()
	p.FirstPageURL = p.first()

	tx = db.Preload(clause.Associations).Limit(p.PerPage).Offset(p.From).Find(p.Data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return p, nil
}

func (p *Paginator) next() string {
	if p.LastPage > p.CurrentPage+1 {
		v := url.Values{}
		p.perPage(v)
		page(v, p.CurrentPage+1)
		return p.Path + "?" + v.Encode()
	}
	return ""
}

func (p *Paginator) prev() string {
	if p.CurrentPage-1 > 0 {
		v := url.Values{}
		p.perPage(v)
		page(v, p.CurrentPage-1)
		return p.Path + "?" + v.Encode()
	}
	return ""
}

func (p *Paginator) last() string {
	if p.LastPage > 0 {
		v := url.Values{}
		p.perPage(v)
		page(v, p.LastPage)
		return p.Path + "?" + v.Encode()
	}
	return ""
}
func (p *Paginator) first() string {
	if p.Total > 0 {
		v := url.Values{}
		p.perPage(v)
		page(v, 1)
		return p.Path + "?" + v.Encode()
	}
	return ""
}

func page(values url.Values, i int) {
	values.Set("page", strconv.Itoa(i))
}

func perPage(values url.Values, i int) {
	values.Set("per_page", strconv.Itoa(i))
}

func (p *Paginator) page(values url.Values, i int) {
	values.Set("per_page", strconv.Itoa(i))
}

func (p *Paginator) perPage(values url.Values) {
	if p.PerPage != DefaultPaginatorPerPage {
		perPage(values, p.PerPage)
	}
}
