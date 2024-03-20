package roulette

import (
	"github.com/jamieyoung5/pooblet/internal/osm/overpass"
	"math/rand"
)

type Game struct{}

var amenities []string = []string{"pub", "bar"}

func NewGame() *Game {
	return &Game{}
}

func (game *Game) Play(lat, long string, radius string) (*Pub, error) {
	var results []overpass.Places
	for _, amenity := range amenities {
		result, err := overpass.GetAmenitiesInRadius(lat, long, radius, amenity)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	places := game.combinePlaces(results)
	randomPlace := places[game.getRandomPlace(places)]

	return parsePlaceToPub(randomPlace)
}

func (game *Game) getRandomPlace(places overpass.Places) int {
	placeIndex := rand.Intn(len(places))
	for id, _ := range places {
		if placeIndex == 0 {
			return id
		}
		placeIndex--
	}
	panic("unreachable")
}

func (game *Game) combinePlaces(placesByAmenity []overpass.Places) (combinedPlaces overpass.Places) {
	combinedPlaces = make(overpass.Places)

	for _, places := range placesByAmenity {
		for id, place := range places {
			if _, ok := combinedPlaces[id]; !ok {
				combinedPlaces[id] = place
			}
		}
	}

	return combinedPlaces
}
