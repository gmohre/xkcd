package api

import (
	"encoding/json"
	"net/http"

	"github.com/gmohre/xkcd/models"
)

func (api *API) CreateFavorite(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	jsondata := models.Favorite{}
	err := decoder.Decode(&jsondata)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
	}
	api.favorites.AddFavorite(jsondata.Number)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("OK"))
}

func (api *API) GetFavorites(w http.ResponseWriter, r *http.Request) {
	favorites := api.favorites.GetFavorites()
	js, _ := json.Marshal(favorites)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
