package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"server/pkg/config"
)

// DB 对象
var DB *gorm.DB
var SqLDB *sql.DB

// Connect  连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}
	// 获取底层的 sqlDB
	SqLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {

	var err error

	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMysqlDatabase()
	case "sqlite":
		deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteAllSqliteTables() error {
	var tables []string
	DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'")
	for _, table := range tables {
		DB.Migrator().DropTable(table)
	}
	return nil
}

func deleteMysqlDatabase() error {
	dbname := CurrentDatabase()
	sql := fmt.Sprintf("DROP DATABASE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("CREATE DATABASE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("USE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}
