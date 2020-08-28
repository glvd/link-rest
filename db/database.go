package db

import (
	"fmt"
	"net/url"

	"github.com/goextension/extmap"
	"github.com/xormsharp/xorm"
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
	ConnectParams() (string, string)
}

func ParseFromMap(m extmap.Map) *SQLConnect {
	c := defaultSQLConnect()
	if m == nil {
		return c
	}
	c.SQLType = m.GetStringD("SQLType", c.SQLType)
	c.Username = m.GetStringD("Username", c.Username)
	c.Password = m.GetStringD("Password", c.Password)
	c.Addr = m.GetStringD("Addr", c.Addr)
	c.Port = m.GetStringD("Port", c.Port)
	c.Schema = m.GetStringD("Schema", c.Schema)
	c.Param = m.GetStringD("Param", c.Param)
	c.Location = m.GetStringD("Location", c.Location)
	return c
}

func defaultSQLConnect() *SQLConnect {
	return &SQLConnect{
		SQLType:  "mysql",
		Username: "root",
		Password: "111111",
		Addr:     "127.0.0.1",
		Port:     "3306",
		Schema:   "linker",
		Param:    "?",
		Location: "Asia/Shanghai",
	}
}

func New(c Connectable) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(c.ConnectParams())
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

func (c SQLConnect) ConnectParams() (string, string) {
	return c.Type(), c.String()
}
