package api

import "github.com/gmohre/xkcd/models"

type API struct {
	comics    *models.ComicManager
	favorites *models.FavoriteManager
}

func NewAPI(db *models.DB) *API {

	comicmgr, _ := models.NewComicManager(db)
	favmgr, _ := models.NewFavoriteManager(db)

	return &API{
		comics:    comicmgr,
		favorites: favmgr,
	}
}
