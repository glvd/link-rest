package model

import (
	"fmt"
	"github.com/xormsharp/xorm"
	"reflect"
)

var migrateTable = map[string]interface{}{}

func RegisterTable(v interface{}) {
	migrateTable[reflect.TypeOf(v).String()] = v
}

func Migration(db *xorm.Engine) {
	for t, tb := range migrateTable {
		ret := db.Sync2(tb)
		if ret != nil {
			fmt.Printf("sync(%v) failed with error:%v\n", t, ret.Error)
			continue
		}
	}
}
