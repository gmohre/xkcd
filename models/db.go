package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	*gorm.DB
}

func NewSqliteDB(databaseName string) *DB {

	db, err := gorm.Open("sqlite3", databaseName)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	//db.LogMode(true)

	return &DB{db}
}
