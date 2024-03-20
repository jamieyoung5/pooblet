// An interface to the overpass API and the queries served to it

package overpass

import (
	"encoding/json"
	"fmt"
)

func GetAmenitiesInRadius(lat, long, radius string, amenity string) (amenitiesInRadius Places, err error) {
	locationRadiusParameter := fmt.Sprintf("(around:%s,%s,%s);", radius, lat, long)
	query := `[out:json];
    (
      node["amenity"="` + amenity + `"]` + locationRadiusParameter + `
      way["amenity"="` + amenity + `"]` + locationRadiusParameter + `
      relation["amenity"="` + amenity + `"]` + locationRadiusParameter + `
    );
    out body;
    >;
    out skel qt;`

	response, err := Query(query)
	if err != nil {
		return nil, err
	}

	var parsedResponse *Response
	if err = json.Unmarshal(response, &parsedResponse); err != nil {
		return nil, err
	}

	return mapResponseElementToId(parsedResponse), nil
}

func mapResponseElementToId(response *Response) (place Places) {
	place = make(Places)
	for _, element := range response.Elements {
		place[element.ID] = element
	}

	return place
}
