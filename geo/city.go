package geo

import "fmt"

type City struct {
	Location
	name        string
	inhabitants int
}

var GermanCities = []City{
	City{Location{50.7273938, 12.4169434}, "Zwickau", 91000},
	City{Location{52.5076274, 13.1442835}, "Berlin", 3470000},
	City{Location{53.5586937, 9.7873987}, "Hamburg", 1763000},
	City{Location{48.1550542, 11.4014126}, "München", 1430000},
	City{Location{50.9573771, 6.8268994}, "Köln", 1047000},
	City{Location{50.1213475, 8.4961409}, "Frankfurt/Main", 718000},
	City{Location{48.77935, 9.0367922}, "Stuttgart", 612000},
	City{Location{51.2385411, 6.7441397}, "Düsseldorf", 605000},
	City{Location{51.5080144, 7.3298396}, "Dortmund", 581000},
	City{Location{51.4409932, 6.9458078}, "Essen", 574000},
	City{Location{53.1203388, 8.4554331}, "Bremen", 552000},
	City{Location{51.3419129, 12.2532121}, "Leipzig", 544000},
	City{Location{51.0769653, 13.6321627}, "Dresden", 536000},
	City{Location{52.3796457, 9.6914329}, "Hannover", 524000},
	City{Location{49.7780678, 9.9080095}, "Würzburg", 125000},

}

var Zwickau = GermanCities[0]

//Methode für Print
func (c City) String() string {
	return fmt.Sprintf("%v@[%v, %v](%v)", c.name, c.lat, c.lon, c.inhabitants)
}

var CityDistanceTable = makeDistanceTable()

func makeDistanceTable() (dt map[int]float64) {
	dt = make(map[int]float64)
	var from, to int
	for from = 0; from < len(GermanCities) - 1; from++ {
		for to = from + 1; to < len(GermanCities); to++ {
			val := GermanCities[from].Dist(GermanCities[to])
			dt[from << 8 + to] = val
			dt[from + to << 8] = val
		}
	}
	return dt
}