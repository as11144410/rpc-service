package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rpc-service/internal/pkg/models"
	"rpc-service/internal/pkg/utils"
)

var Db *sql.DB

// create database if not exists shopify_checkout character set utf8mb4 collate utf8mb4_bin;
func init() {
	InstallDb()
}

func InstallDb() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/shopify_checkout?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetEvnWithDefaultVal("MYSQL_USER", "root"), utils.GetEvnWithDefaultVal("MYSQL_PASSWORD", "123456"), utils.GetEvnWithDefaultVal("MYSQL_ADDR", "localhost:3306"))
	cnn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Db, _ = cnn.DB()

	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(100)

	cnn.AutoMigrate(&models.User{})
}

func GetDb() *sql.DB {
	return Db
}
