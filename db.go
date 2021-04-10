package db

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

/*
 * #add mysql user
 * grant all privileges on *.* to 'sysuser'@'localhost' identified by 'sysdba' with grant option;
 * flush privileges;
 *
 * @func db connect
 */
func connect(alias string) (db *gorm.DB, err error) {
	Configuration()
	cfg := defaultDbConfig(conf.Database[alias])
	dsn := fmt.Sprintf(
		"%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Dsn,
	)
	fmt.Println("db init, cfg:", cfg)

	var ormLogger logger.Interface
	if cfg.Debug {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	ormCfg := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 ormLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	}

	db, err = gorm.Open(mysql.Open(dsn), ormCfg)
	if err != nil {
		os.Exit(0)
		return nil, err
	}
	conn, err := db.DB()
	conn.SetMaxIdleConns(cfg.MaxIdleConns)
	conn.SetMaxOpenConns(cfg.MaxOpenConns)
	conn.SetConnMaxLifetime(time.Hour * 8)
	return db, err
}

const (
	DEFAULT = "db"
)

/*
 * @func set db
 */
func New(dbname string) (db *gorm.DB) {
	db, _ = dataBaseCache.get(dbname)
	if db == nil {
		newdb, _ := connect(dbname)
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

// database alias cache.
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
