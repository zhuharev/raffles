package models

import (
	"database/sql"
	"fmt"
	//"os"
	//"path/filepath"
	"reflect"

	"github.com/go-xorm/xorm"
	//_ "github.com/mattn/go-sqlite3"
	//"github.com/ungerik/go-dry"
	_ "github.com/go-xorm/ql"
	_ "github.com/lunny/ql/driver"

	"modules/setting"
)

var (
	x      *xorm.Engine
	tables = []interface{}{
		new(User),
		new(Likes),
		new(Source),
		new(Task)}
)

// Engine represents a xorm engine or session.
type Engine interface {
	Delete(interface{}) (int64, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Id(interface{}) *xorm.Session
	Sql(string, ...interface{}) *xorm.Session
	Where(string, ...interface{}) *xorm.Session
}

func NewEngine() (e error) {
	x, e = xorm.NewEngine(setting.Config.Db.Driver, setting.Config.Db.Setting)
	if e != nil {
		return e
	}

	/*logPath := filepath.Join(setting.LogDir, "sql.log")
	f, e := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		panic(e)
	}
	logger := xorm.NewSimpleLogger(f)
	x.SetLogger(logger)*/
	x.ShowSQL = true
	x.ShowDebug = true

	e = x.Sync2(tables...)
	if e != nil {
		fmt.Println(e)
	}

	return nil
}

func Delete(bean interface{}) (e error) {
	_, e = x.Delete(bean)
	return
}

func Save(bean interface{}) (e error) {
	val := reflect.ValueOf(bean)
	val = val.Elem()
	idVal := val.FieldByName("Id")
	if !idVal.IsValid() {
		return fmt.Errorf("id value is nil")
	}
	id := idVal.Interface().(int64)

	if id == 0 {
		_, e = x.InsertOne(bean)
		return
	} else {
		_, e = x.Id(id).Update(bean)
		return
	}
	return
}
