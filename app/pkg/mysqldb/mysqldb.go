package mysqldb

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yanggelinux/cattle/global"
	"github.com/yanggelinux/cattle/pkg/log"
	"go.uber.org/zap"
	"time"
)

type MySQLDB struct {
	db          *sql.DB
	DSN         string
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}

func (m *MySQLDB) MySQLOpen() error {
	db, err := sql.Open("mysql", m.DSN)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(m.Active)
	db.SetMaxIdleConns(m.Idle)
	db.SetConnMaxLifetime(time.Duration(m.IdleTimeout))
	m.db = db
	return nil
}

func (m *MySQLDB) MySQLClose() {
	err := m.db.Close()
	if err != nil {
		log.Logger.Error("close mysql error:", zap.Error(err))
		return
	}
}

// select 查询数据
func (m *MySQLDB) MySQLQuery(sql string) (*sql.Rows, error) {
	rows, err := m.db.Query(sql)
	return rows, err
}
func (m *MySQLDB) MySQLQueryContext(ctx context.Context, sql string) (*sql.Rows, error) {
	rows, err := m.db.QueryContext(ctx, sql)
	return rows, err
}

func NewMySQLDB(user, password, host string, port int, dbName string) *MySQLDB {
	var mysqlDB *MySQLDB = new(MySQLDB)
	mysqlDB.DSN = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, password, host, port, dbName)
	mysqlDB.Active = global.MySQLSetting.MaxOpenConns
	mysqlDB.Idle = global.MySQLSetting.MaxIdleConns
	mysqlDB.IdleTimeout = time.Duration(global.MySQLSetting.ConnMaxLifetime) * time.Hour
	return mysqlDB
}
