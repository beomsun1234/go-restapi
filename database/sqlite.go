package database

import (
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Sqlite struct {
	SqliteDB *gorm.DB
}

// github.com/mattn/go-sqlite3

func NewSqlite() *Sqlite {
	return &Sqlite{}
}

func (c *Sqlite) Connection() {
	var err error
	c.SqliteDB, err = gorm.Open(sqlite.Open("../database/sqlitedb/gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
