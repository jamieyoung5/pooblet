package main

import (
	"fmt"
	"github.com/jamieyoung5/pooblet/internal/osm/overpass"
	"github.com/jamieyoung5/pooblet/internal/roulette"
	"github.com/jamieyoung5/pooblet/internal/verification"
	"sort"
)

// Function to execute the Overpass query
func getPubs(lat, lon float64, rad int) {
	latitude, longitude, err := verification.VerifyLocation(lon, lat)
	if err != nil {
		panic(err)
	}

	radius, err := verification.VerifyRadius(rad)
	if err != nil {
		panic(err)
	}

	response, err := overpass.GetAmenitiesInRadius(latitude, longitude, radius, "bar")
	if err != nil {
		panic(err)
	}

	tags := make(map[string]int)
	for _, element := range response {
		for tag, _ := range element.Tags {
			tags[tag] += 1
		}
	}

	keys := make([]string, 0, len(tags))

	for key := range tags {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return tags[keys[i]] > tags[keys[j]]
	})

	for _, k := range keys {
		msg := fmt.Sprintf("%s: %d", k, tags[k])
		fmt.Println(msg)
	}

}

func getPubsForRealzies(lat, lon float64, rad int) {
	pubRoulette := roulette.NewGame()
	latitude, longitude, err := verification.VerifyLocation(lon, lat)
	if err != nil {
		panic(err)
	}
	radius, err := verification.VerifyRadius(rad)
	if err != nil {
		panic(err)
	}

	pub, err := pubRoulette.Play(latitude, longitude, radius)
	if err != nil {
		panic(err)
	}

	fmt.Println(pub)
}

func main() {
	lat, lon := 55.953251, -3.188267
	radius := 500 // in meters

	getPubsForRealzies(lat, lon, radius)
}
