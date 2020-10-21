package db

import (
	"fmt"

	"net/url"

	"github.com/goextension/extmap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
type mysqlInfo struct {
	SQLType  string
	Username string
	Password string
	Addr     string
	Port     string
	Schema   string
	Param    string
	Location string
}

func (m mysqlInfo) Type() string {
	return "mysql"
}

func (m mysqlInfo) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%sloc=%s&charset=utf8&parseTime=true",
		m.Username, m.Password, m.Addr, m.Port, m.Schema, m.Param, url.QueryEscape(m.Location),
	)
}

func (m mysqlInfo) ConnectParams() (string, string) {
	return m.Type(), m.String()
}

type Connectable interface {
	Type() string
	String() string
	ConnectParams() (string, string)
}

func ParseFromMap(m extmap.Map) Connectable {
	if m == nil {
		m = extmap.Map{}
	}
	switch m.GetStringD("SQLType", "mysql") {
	default:
		return parseMysql(m)
	}
}

func parseMysql(m extmap.Map) *mysqlInfo {
	c := &mysqlInfo{
		SQLType:  "mysql",
		Username: "root",
		Password: "111111",
		Addr:     "127.0.0.1",
		Port:     "3306",
		Schema:   "linker",
		Param:    "?",
		Location: "Asia/Shanghai",
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

func New(c Connectable) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.String()), nil)
	if err != nil {
		return nil, fmt.Errorf("connect db error:%w", err)
	}
	return db, nil
}

func OpenSqlite() {
	sqlite.Open("linker.db")
}
