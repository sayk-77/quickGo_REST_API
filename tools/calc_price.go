package tools

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"strconv"
)

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type GeocodeResult struct {
	Length      int          `json:"length"`
	Address     string       `json:"address"`
	Coordinates *Coordinates `json:"coordinates"`
}

func GeocodeByAddresses(addresses []string) (float64, error) {
	results := make([]GeocodeResult, 0)

	for _, address := range addresses {
		resp, err := http.Get("https://nominatim.openstreetmap.org/search?q=" + address + "&format=json")
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()

		var data []map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return 0, err
		}

		var coord *Coordinates
		if len(data) > 0 {
			lat, _ := data[0]["lat"].(string)
			lon, _ := data[0]["lon"].(string)
			coord = &Coordinates{Lat: parseFloat(lat), Lon: parseFloat(lon)}
		}

		results = append(results, GeocodeResult{Address: address, Coordinates: coord})
	}

	if results[0].Coordinates != nil && results[1].Coordinates != nil {
		return calculateDistance(*results[0].Coordinates, *results[1].Coordinates), nil
	}

	return 0, errors.New("не удалось получить координаты для одного из адресов")
}

func calculateDistance(coord1 Coordinates, coord2 Coordinates) float64 {
	const R = 6371

	lat1 := toRadians(coord1.Lat)
	lon1 := toRadians(coord1.Lon)
	lat2 := toRadians(coord2.Lat)
	lon2 := toRadians(coord2.Lon)

	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlon/2)*math.Sin(dlon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
