package db

import (
	"database/sql"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/wjhdec/echo-ext/pkg/config"
	"github.com/wjhdec/echo-ext/pkg/elog"
	"strings"
	"sync"
	"time"
)

var (
	innerDB DB
	once    sync.Once
)

type DB interface {
	GetByName(name string) (*sql.DB, error)
}

func New(cfg config.Config) DB {
	once.Do(func() {
		innerDB = &db{cfg: cfg, dbMap: make(map[string]*sql.DB)}
	})
	return innerDB
}

type db struct {
	dbMap map[string]*sql.DB
	cfg   config.Config
}

func (d *db) GetByName(name string) (db *sql.DB, err error) {
	db = d.dbMap[name]
	if db != nil {
		return
	} else {
		dbCfg := new(Config)
		if err = d.cfg.UnmarshalByKey("db."+name, dbCfg); err != nil {
			return
		}
		if db, err = newDB(dbCfg); err != nil {
			return
		}
		d.dbMap[name] = db
		return
	}
}

func parseQueryLevel(level string) sqldblogger.Level {
	switch strings.ToLower(level) {
	case "trace":
		return sqldblogger.LevelTrace
	case "debug":
		return sqldblogger.LevelDebug
	case "info":
		return sqldblogger.LevelInfo
	case "error":
		return sqldblogger.LevelError
	default:
		return sqldblogger.LevelDebug
	}
}

func newDB(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, cfg.Dsn)
	if err != nil {
		return nil, err
	}
	lvl := parseQueryLevel(cfg.QueryLevel)
	db = sqldblogger.OpenDriver(cfg.Dsn, db.Driver(), NewAdaptor(elog.GlobalLogger()),
		sqldblogger.WithQueryerLevel(lvl),
		sqldblogger.WithPreparerLevel(lvl),
		sqldblogger.WithExecerLevel(lvl))
	db.SetMaxIdleConns(cfg.MaxIdle)
	db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdle) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	return db, nil
}
