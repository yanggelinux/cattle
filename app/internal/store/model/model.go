package model

import (
	"database/sql"
	"fmt"
	"github.com/yanggelinux/cattle/global"
	"github.com/yanggelinux/cattle/pkg/log"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	sqlDB *sql.DB
	err   error
)

type MDB = gorm.DB

var ErrRecordNotFound = gorm.ErrRecordNotFound

type Model struct {
	IsDeleted   int8      `json:"isDeleted" gorm:"default:0;column:is_deleted"`
	DeletedTime time.Time `json:"deletedTime"  gorm:"default:2025-01-01 00:00:00;column:deleted_time"` //自动填充时间到毫秒
	UpdatedTime time.Time `json:"updatedTime"  gorm:"column:updated_time"`
	CreatedTime time.Time `json:"createdTime"  gorm:"column:created_time"`
}

func (m *Model) IsDeletedField() string {
	return "is_deleted"
}
func (m *Model) DeletedTimeField() string {
	return "deleted_time"
}
func (m *Model) UpdatedTimeField() string {
	return "updated_time"
}
func (m *Model) CreatedTimeField() string {
	return "created_time"
}

func updateAndCreateTimeForCreateCallback(DB *gorm.DB) {
	if DB.Statement.Schema != nil {
		//nowTime := util.FormatTimeToString2(time.Now())
		nowTime := time.Now()
		createTimeField := DB.Statement.Schema.LookUpField("CreatedTime")
		if createTimeField != nil {
			DB.Statement.SetColumn("CreatedTime", nowTime, true)
		}

		updateTimeField := DB.Statement.Schema.LookUpField("UpdatedTime")
		if updateTimeField != nil {
			DB.Statement.SetColumn("UpdatedTime", nowTime, true)
		}
	}
}

func updateTimeForUpdateCallback(DB *gorm.DB) {
	if DB.Statement.Schema != nil {
		//nowTime := util.FormatTimeToString2(time.Now())
		nowTime := time.Now()
		updateTimeField := DB.Statement.Schema.LookUpField("UpdatedTime")
		if updateTimeField != nil {
			DB.Statement.SetColumn("UpdatedTime", nowTime, true)

		}

	}
}

func deleteTimeForUpdateCallback(DB *gorm.DB) {
	if DB.Statement.Schema != nil {
		//nowTime := util.FormatTimeToString2(time.Now())
		nowTime := time.Now()
		deleteTimeField := DB.Statement.Schema.LookUpField("DeletedTime")
		if deleteTimeField != nil {
			DB.Statement.SetColumn("DeletedTime", nowTime, true)

		}

	}
}

func RecoverRollback(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
		errMsg := fmt.Sprintf("%s:%+v", "recover panic", r)
		log.Logger.Error(errMsg)
	}
}

func GetDB() *MDB {
	return DB
}

func SetupModel() error {
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=10s&writeTimeout=5s",
		global.MySQLSetting.Username,
		global.MySQLSetting.Password,
		global.MySQLSetting.Host,
		global.MySQLSetting.Port,
		global.MySQLSetting.DBName)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		//执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		//PrepareStmt: true,
	})
	if err != nil {
		log.Logger.Error("connect database error:", zap.Error(err))
		return err
	}

	sqlDB, err = DB.DB()
	if err != nil {
		log.Logger.Error("get db.DB() error:", zap.Error(err))
		return err
	}
	// callback
	err = DB.Callback().Create().Before("gorm:create").Register("createUpdateTime", updateAndCreateTimeForCreateCallback)
	if err != nil {
		log.Logger.Error("updateAndCreateTimeForCreateCallback Register error:", zap.Error(err))
		return err
	}
	err = DB.Callback().Update().Before("gorm:update").Register("updateTime", updateTimeForUpdateCallback)
	if err != nil {
		log.Logger.Error("updateTimeForUpdateCallback Register error:", zap.Error(err))
		return err
	}
	err = DB.Callback().Delete().Before("gorm:delete").Register("deleteTime", deleteTimeForUpdateCallback)
	if err != nil {
		log.Logger.Error("deleteTimeForDeleteCallback Register error:", zap.Error(err))
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(global.MySQLSetting.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(global.MySQLSetting.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(global.MySQLSetting.ConnMaxLifetime) * time.Hour)

	//创建表 建议手动创建表
	//err = DB.AutoMigrate(&User{}, &Role{})
	//if err != nil {
	//	log.Logger.Error("create db table error:", zap.Error(err))
	//	return
	//}
	return nil
}
