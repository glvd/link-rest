package db

import (
	"fmt"
	"github.com/xormsharp/xorm"
	"net/url"
)

type SQLConnect struct {
	SQLType  string
	Username string
	Password string
	Addr     string
	Port     string
	Schema   string
	Param    string
	Location string
}

type Connectable interface {
	Type() string
	String() string
}

func defaultConnectAble() Connectable {
	return &SQLConnect{
		SQLType:  "mysql",
		Username: "root",
		Password: "111111",
		Addr:     "127.0.0.1",
		Port:     "3306",
		Schema:   "redball",
		Param:    "?",
		Location: "Asia/Shanghai",
	}
}

func New(c Connectable) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(c.Type(), c.String())
	if err != nil {
		return nil, err
	}
	return engine, nil
}

func (c SQLConnect) Type() string {
	return c.SQLType
}

func (c SQLConnect) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%sloc=%s&charset=utf8&parseTime=true",
		c.Username, c.Password, c.Addr, c.Port, c.Schema, c.Param, url.QueryEscape(c.Location),
	)
}
