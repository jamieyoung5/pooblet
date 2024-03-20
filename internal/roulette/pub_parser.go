package roulette

import (
	"github.com/jamieyoung5/pooblet/internal/osm/nominatim"
	"github.com/jamieyoung5/pooblet/internal/osm/overpass"
	"github.com/jamieyoung5/pooblet/internal/osm/overpass/tags_filter"
)

func parsePlaceToPub(place overpass.Element) (*Pub, error) {
	names := tags_filter.FilterPlaceNameFromTags(place.Tags)
	tags := tags_filter.FilterTags(place.Tags, tags_filter.ValidTags)
	address, err := nominatim.ReverseGeocode(place.Lat, place.Lon)
	if err != nil {
		return nil, err
	}

	return &Pub{
		Tags:      tags,
		Longitude: place.Lon,
		Latitude:  place.Lat,
		Address:   address,
		Name:      names,
	}, nil
}
