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

func Migration(db *xorm.Engine) (err error) {
	for t, tb := range migrateTable {
		err = db.Sync2(tb)
		if err != nil {
			fmt.Printf("sync(%v) failed with error:%v\n", t, err)
			return err
		}
	}
	return nil
}
