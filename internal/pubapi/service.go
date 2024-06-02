package pubapi

import (
	"github.com/jamieyoung5/pooblet/internal/roulette"
	"github.com/jamieyoung5/pooblet/internal/verification"
)

func GetPub(lat, lon float64, rad int) (roulette.Pub, error) {
	latitude, longitude, err := verification.VerifyLocation(lon, lat)
	if err != nil {
		return roulette.Pub{}, err
	}
	radius, err := verification.VerifyRadius(rad)
	if err != nil {
		return roulette.Pub{}, err
	}

	pub, err := roulette.NewGame().Play(latitude, longitude, radius)
	if err != nil {
		return roulette.Pub{}, err
	}

	return *pub, nil
}
