package db

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xf005/goini"
	"github.com/xf005/logger"
)

const (
	DEFAULT = "db"
)

/*
 * add mysql user
 * #grant all privileges on *.* to 'sysuser'@'localhost' identified by 'sysdba' with grant option;
 * #flush privileges;
 *
 * @func init db connect
 */
func newDB(dbname string) (db *gorm.DB, err error) {
	//conf := goini.SetConfig("C:/server/goworkspace/workspace/src/user/conf/app.conf")
	conf := goini.SetConfig("./conf/app.conf")
	host := conf.GetValue(dbname, "host")
	port := conf.GetValue(dbname, "port")
	name := conf.GetValue(dbname, "name")
	user := conf.GetValue(dbname, "user")
	pass := conf.GetValue(dbname, "pass")
	mode := conf.GetValue(dbname, "mode")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)
	fmt.Println(dns)
	db, err = gorm.Open("mysql", dns)
	//db, err := gorm.Open("mysql", "sysuser:syspass@tcp(localhost:3306)/adshared?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error("Database %s init error. %s", dbname, err.Error())
		os.Exit(0)
		return nil, err
	}
	logger.Info("Database %s init.", dbname)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)
	db.DB().SetConnMaxLifetime(time.Hour * 7)
	db.SingularTable(true)
	//db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&User{})
	db.LogMode(mode=="true")
	return db, nil
}

/*
 * @func set db
 */
func NewDB(dbname string) (db *gorm.DB) {
	db, _ = dataBaseCache.get(dbname)
	if db == nil {
		newdb, _ := newDB(dbname)
		db = newdb
		dataBaseCache.add(dbname, db)
	}
	return db
}

/*
 * @func default db
 */
func DB() (db *gorm.DB) {
	return NewDB(DEFAULT)
}

var (
	dataBaseCache = &dbCache{cache: make(map[string]*gorm.DB)}
)

// database alias cacher.
type dbCache struct {
	mux   sync.RWMutex
	cache map[string]*gorm.DB
}

// add database alias with original name.
func (ac *dbCache) add(name string, al *gorm.DB) (added bool) {
	ac.mux.Lock()
	defer ac.mux.Unlock()
	if _, ok := ac.cache[name]; !ok {
		ac.cache[name] = al
		added = true
	}
	return
}

// get database alias if cached.
func (ac *dbCache) get(name string) (al *gorm.DB, ok bool) {
	ac.mux.RLock()
	defer ac.mux.RUnlock()
	al, ok = ac.cache[name]
	return
}
