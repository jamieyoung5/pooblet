package roulette

import (
	"github.com/jamieyoung5/pooblet/internal/osm/nominatim"
	"github.com/jamieyoung5/pooblet/internal/osm/overpass/tags_filter"
)

type Pub struct {
	Tags      []string
	Longitude float64
	Latitude  float64
	Address   *nominatim.Address
	Name      tags_filter.Names
}

var amenities = []string{"pub", "bar"}
