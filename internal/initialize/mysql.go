package initialize

import (
	"fmt"
	"time"

	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	mysqlConfig := global.Config.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("failed to connect database :: " + err.Error())
	}

	global.MDB = db

	// SetPool
	SetPool()

	// migrateTables
	migrateTables()
}

func SetPool() {
	mysqlConfig := global.Config.Mysql
	sqlDB, err := global.MDB.DB()
	if err != nil {
		panic("failed to connect database :: " + err.Error())
	}

	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConfig.MaxOpen))

}

func migrateTables() {
	err := global.MDB.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.GroupUser{},
	)
	if err != nil {
		panic("failed to migrate tables :: " + err.Error())
	}
}
