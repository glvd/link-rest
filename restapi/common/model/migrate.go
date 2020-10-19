package model

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
)

var migrateTable = map[string]interface{}{}

func RegisterTable(v interface{}) {
	if t, b := v.(schema.Tabler); b {
		migrateTable[t.TableName()] = v
		return
	}
	migrateTable[reflect.TypeOf(v).String()] = v
}

func Migration(db *gorm.DB) (err error) {
	for t, tb := range migrateTable {
		fmt.Println("migrate table", t)
		err = db.AutoMigrate(tb)
		if err != nil {
			fmt.Printf("sync(%v) failed with error:%v\n", t, err)
			return err
		}
	}
	return nil
}
