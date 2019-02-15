package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Favorite struct {
	Number int `json:"num"`
}

type Favorites []Favorite

type FavoriteManager struct {
	db *DB
}

func (state *FavoriteManager) AddFavorite(number int) *Favorite {
	favorite := &Favorite{
		Number: number,
	}
	state.db.Create(&favorite)
	return favorite
}

func (state *FavoriteManager) GetFavorites() *Favorites {
	favorites := Favorites{}
	state.db.Find(&favorites)
	return &favorites
}

func NewFavoriteManager(db *DB) (*FavoriteManager, error) {

	db.AutoMigrate(&Favorite{})
	mgr := FavoriteManager{}
	mgr.db = db
	return &mgr, nil
}
