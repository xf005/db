package db

import (
	"fmt"
	"github.com/xf005/db/conf"
	"gorm.io/gorm/schema"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DEFAULT = "db"
)

/*
 * #add mysql user
 * grant all privileges on *.* to 'sysuser'@'localhost' identified by 'sysdba' with grant option;
 * flush privileges;
 *
 * @func init db connect
 */
func newDB(dbname string) (db *gorm.DB, err error) {
	app := conf.NewConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", app.Db.User, app.Db.Pass, app.Db.Host, app.Db.Port, app.Db.Name)
	gormConf := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	}
	if app.Db.LogMode {
		gormConf.Logger = logger.Default.LogMode(logger.Info)
	}
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), gormConf)
	if err != nil {
		os.Exit(0)
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour * 8)
	return db, err
}

/*
 * @func set db
 */
func New(dbname string) (db *gorm.DB) {
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
	return New(DEFAULT)
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
