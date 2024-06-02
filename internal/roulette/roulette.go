package roulette

import (
	"errors"
	"github.com/jamieyoung5/pooblet/internal/osm/overpass"
	"math/rand"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Play(lat, long string, radius string) (*Pub, error) {
	var results []overpass.Places
	for _, amenity := range amenities {
		result, err := overpass.GetAmenitiesInRadius(lat, long, radius, amenity)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	places := g.combinePlaces(results)
	if len(places) <= 0 {
		return nil, errors.New("no places found")
	}
	randomPlace := places[g.getRandomPlace(places)]

	return parsePlaceToPub(randomPlace)
}

func (g *Game) getRandomPlace(places overpass.Places) int {
	placeIndex := rand.Intn(len(places))
	for id, _ := range places {
		if placeIndex == 0 {
			return id
		}
		placeIndex--
	}

	return 0
}

func (g *Game) combinePlaces(placesByAmenity []overpass.Places) (combinedPlaces overpass.Places) {
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
