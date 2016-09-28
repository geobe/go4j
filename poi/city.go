package poi

import (
	"fmt"
	"github.com/geobe/go4j/loc"
)

type City struct {
	loc.Location
	name        string
	inhabitants int
}

func (c City) Name() string {
	return c.name
}

func (c City) Inhabitants() int {
	return c.inhabitants
}

var GermanCities = []City{
	// Literal mit benannten Strukturkomponenten anlegen
	{Location: loc.New(50.7273938, 12.4169434), name: "Zwickau", inhabitants: 91000},
	// Literal mit positionsabhängigen Strukturkomponenten anlegen
	{loc.New(52.5076274, 13.1442835), "Berlin", 3470000},
	{loc.New(53.5586937, 9.7873987), "Hamburg", 1763000},
	{loc.New(48.1550542, 11.4014126), "München", 1430000},
	{loc.New(50.9573771, 6.8268994), "Köln", 1047000},
	{loc.New(50.1213475, 8.4961409), "Frankfurt/Main", 718000},
	{loc.New(48.77935, 9.0367922), "Stuttgart", 612000},
	{loc.New(51.2385411, 6.7441397), "Düsseldorf", 605000},
	{loc.New(51.5080144, 7.3298396), "Dortmund", 581000},
	{loc.New(51.4409932, 6.9458078), "Essen", 574000},
	{loc.New(53.1203388, 8.4554331), "Bremen", 552000},
	{loc.New(51.3419129, 12.2532121), "Leipzig", 544000},
	{loc.New(51.0769653, 13.6321627), "Dresden", 536000},
	{loc.New(52.3796457, 9.6914329), "Hannover", 524000},
	{loc.New(49.7780678, 9.9080095), "Würzburg", 125000},
}

var Zwickau = GermanCities[0]

//Methode für Print
func (c City) String() string {
	lat, lon := c.LatLon()
	return fmt.Sprintf("%v@[%v, %v](%v)", c.name, lat, lon, c.inhabitants)
}

var CityDistanceTable = makeDistanceTable()

func makeDistanceTable() (dt map[int]float64) {
	dt = make(map[int]float64)
	var from, to int
	for from = 0; from < len(GermanCities)-1; from++ {
		for to = from + 1; to < len(GermanCities); to++ {
			val := GermanCities[from].Dist(GermanCities[to])
			dt[from<<8+to] = val
			dt[from+to<<8] = val
		}
	}
	return dt
}
