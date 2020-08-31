package model

import (
	"github.com/goextension/log"
	"github.com/xormsharp/xorm"
	"math"
	"net/http"
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

type Counter interface {
	Count(session *xorm.Session) (int64, error)
}

func defaultPage(v interface{}) *Paginator {
	return &Paginator{
		CurrentPage: 1,
		Data:        v,
	}
}

func Page(v interface{}, r *http.Request) *Paginator {
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
	//last := urls.Get("last_page")
	//p.LastPage, err = strconv.Atoi(last)
	//if err != nil {
	//	p.LastPage = int(math.Ceil(float64(p.Total)/float64(p.PerPage))) - p.CurrentPage
	//}
	//from := urls.Get("from")
	//p.From, err = strconv.Atoi(from)
	//if err != nil {
	//	p.From = (p.CurrentPage - 1) * p.PerPage
	//}
	//
	//to := urls.Get("to")
	//p.To, err = strconv.Atoi(to)
	//if err != nil {
	//	p.To = p.From + p.PerPage
	//}

	return p
}

func (p *Paginator) Find(session *xorm.Session) (*Paginator, error) {
	count, err := session.Count()
	log.Infow("count", "error", err, "count", count)
	if err != nil {
		return nil, err
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

	err = session.Limit(p.PerPage, p.From).Find(p.Data)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Paginator) next() string {
	if p.LastPage > p.CurrentPage+1 {
		return p.Path + "?" + page(p.CurrentPage+1)
	}
	return ""
}

func (p *Paginator) prev() string {
	if p.CurrentPage-1 > 0 {
		return p.Path + "?" + page(p.CurrentPage-1)
	}
	return ""
}

func (p *Paginator) last() string {
	if p.LastPage > 0 {
		return p.Path + "?" + page(p.LastPage)
	}
	return ""
}
func (p *Paginator) first() string {
	if p.Total > 0 {
		return p.Path + "?" + page(1)
	}
	return ""
}

func page(i int) string {
	return "page=" + strconv.Itoa(i)
}
