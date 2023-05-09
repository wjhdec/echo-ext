package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/wjhdec/echo-ext/pkg/config"
	"github.com/wjhdec/echo-ext/pkg/elog"
	"sync"
	"time"
)

var (
	innerDB DB
	once    sync.Once
)

type DB interface {
	GetByName(name string) (*sqlx.DB, error)
}

func New(cfg config.Config) DB {
	once.Do(func() {
		innerDB = &db{cfg: cfg, dbMap: make(map[string]*sqlx.DB)}
	})
	return innerDB
}

type db struct {
	dbMap map[string]*sqlx.DB
	cfg   config.Config
}

func (d *db) GetByName(name string) (*sqlx.DB, error) {
	db := d.dbMap[name]
	if db != nil {
		return db, nil
	} else {
		dbCfg := new(Config)
		err := d.cfg.UnmarshalByKey("db."+name, dbCfg)
		if err != nil {
			return nil, err
		}
		db, err := newDB(dbCfg)
		if err != nil {
			return nil, err
		}
		d.dbMap[name] = db
		return db, nil
	}
}

func newDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sql.Open(cfg.Driver, cfg.Dsn)
	if err != nil {
		return nil, err
	}
	db = sqldblogger.OpenDriver(cfg.Dsn, db.Driver(), NewAdaptor(elog.GlobalLogger()))
	connect := sqlx.NewDb(db, cfg.Driver)
	connect.SetMaxIdleConns(cfg.MaxIdle)
	connect.SetConnMaxIdleTime(time.Duration(cfg.MaxIdle) * time.Second)
	connect.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	return connect, nil
}
