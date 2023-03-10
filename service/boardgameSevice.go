package service

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/broemp/brettspielTrackerAPI/entity"
	"github.com/broemp/brettspielTrackerAPI/initializers"
)

const BGG_API_URL = "https://boardgamegeek.com/xmlapi2/"

type BoardgameService interface {
	RandomBoardgameFromBGGCollection(username string) (entity.Boardgame, error)
}

type boardgameService struct {
}

func NewBoardgameRepository() BoardgameService {
	return &boardgameService{}
}

// Get Random Boardgame from local Database or if no
// Data is found from BGG API
func (service *boardgameService) RandomBoardgameFromBGGCollection(username string) (entity.Boardgame, error) {

	var collection entity.Collection
	initializers.DB.Preload("Games").First(&collection, "username = ?", username)

	if collection.Username != username {
		collection = pullCollection(username)
		initializers.DB.Create(&collection)
	} else {
	}

	length := len(collection.Games)

	if length == 0 {
		return entity.Boardgame{}, errors.New("Empty Collection")
	}

	randomBGGID := collection.Games[rand.Intn(length)].BGGID

	var boardgame entity.Boardgame
	initializers.DB.First(&boardgame, "bgg_id = ?", randomBGGID)

	return boardgame, nil
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		println("Failed to convert String: ", err)
	}

	return i
}

// Pull Collection from BGG
// Create not fully filled Game objects for collection
func pullCollection(username string) (collection entity.Collection) {
	apiQuery := "https://boardgamegeek.com/xmlapi2/collection?username=QUERY&subtype=boardgame&stats=1"
	apiQuery = strings.Replace(apiQuery, "QUERY", username, 1)
	counter := 1

request:
	// HTTP Request
	resp, err := http.Get(apiQuery)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// Parse xml and check for error
	var bggCollection entity.BGGCollection
	var bggError entity.Message
	xml.Unmarshal(body, &bggError)

	if bggError.Message != "" {
		println("Repeat Request: Wait time ", counter, " Seconds")
		time.Sleep(time.Duration(counter) * time.Second)
		counter *= 2
		goto request
	}

	xml.Unmarshal(body, &bggCollection)

	collection = entity.Collection{
		Username: username,
	}
	for _, bggThing := range bggCollection.Item {
		var boardgame entity.Boardgame
		initializers.DB.First(&boardgame, "bgg_id = ?", bggThing.Objectid)

		if boardgame.BGGID == stringToInt(bggThing.Objectid) && bggThing.Status.Own == "1" {
			collection.Games = append(collection.Games, boardgame)
		} else if bggThing.Status.Own == "1" {
			rating, err := strconv.ParseFloat(bggThing.Stats.Rating.Average.Value, 8)
			println(rating)

			if err != nil {
				println(err)
			}

			boardgame = entity.Boardgame{
				Name:        bggThing.Name.Text,
				BGGID:       stringToInt(bggThing.Objectid),
				ImageUrl:    bggThing.Image,
				ReleaseYear: bggThing.Yearpublished,
				Rating:      rating,
				MinPlayer:   stringToInt(bggThing.Stats.Minplayers),
				MaxPlayer:   stringToInt(bggThing.Stats.Maxplayers),
				MinPlaytime: stringToInt(bggThing.Stats.Minplaytime),
				MaxPlaytime: stringToInt(bggThing.Stats.Maxplaytime),
			}

			initializers.DB.Create(&boardgame)
			collection.Games = append(collection.Games, boardgame)
		}

	}

	return
}
