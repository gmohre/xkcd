package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Comic struct {
	Number int    `json:"num"`
	Title  string `json:"title"`
	Image  string `json:"img"`
	Year   string `json:"year"`
	Month  string `json:"month"`
	Day    string `json:"day"`
}

type Comics []Comic

type ComicManager struct {
	db *DB
}

func NewComicManager(db *DB) (*ComicManager, error) {

	db.AutoMigrate(&Comic{})
	mgr := ComicManager{}
	mgr.db = db
	return &mgr, nil
}
