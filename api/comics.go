package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gmohre/xkcd/models"
	"github.com/gorilla/mux"
)

func (api *API) GetLatestComic(w http.ResponseWriter, r *http.Request) {
	comic_url := "http://xkcd.now.sh/"
	newestComic := getComicByURL(comic_url)

	fmt.Println(newestComic.Title)
	fmt.Println(newestComic.Number)
	back := r.URL.Query().Get("backissues")

	comics := models.Comics{
		newestComic,
	}
	if back != "" {
		start := newestComic.Number - 1
		fmt.Println(start)
		back, _ := strconv.ParseInt(back, 10, 32)
		end := start - int(back)
		for i := start; i > end; i-- {
			comic_url := "http://xkcd.now.sh/" + strconv.Itoa(i)
			fmt.Println(comic_url)
			comic := getComicByURL(comic_url)
			comics = append(comics, comic)
		}
	}
	json.NewEncoder(w).Encode(comics)
}

func (api *API) GetComicByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	comicID := vars["comicID"]
	comic_url := "http://xkcd.now.sh/" + comicID
	requestedComic := getComicByURL(comic_url)

	fmt.Println(comicID)
	fmt.Println(requestedComic.Title)
	fmt.Println(requestedComic.Number)

	json.NewEncoder(w).Encode(requestedComic)
}

func getComicByURL(url string) models.Comic {
	client := &http.Client{
		Timeout: time.Second * 20,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkErr(err)

	res, getErr := client.Do(req)
	checkErr(getErr)

	body, readErr := ioutil.ReadAll(res.Body)
	checkErr(readErr)

	requestedComic := models.Comic{}
	jsonErr := json.Unmarshal(body, &requestedComic)
	checkErr(jsonErr)

	return requestedComic
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}
}
