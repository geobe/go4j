package poi

import (
	"fmt"
	"github.com/geobe/go4j/loc"
	"github.com/geobe/go4j/permute"
	"testing"
	"time"
)

// teste, ob Dist auch für City Objekte funktioniert
func TestCities(t *testing.T) {
	for _, city := range GermanCities {
		fmt.Printf("Z -> %s = %.0f km\n",
			city.name, Zwickau.Dist(city))
	}
}

// teste Getter auf City
func TestGetLat(t *testing.T) {
	b := City{loc.New(52.507, 13.144),
		"Berlin", 3470000}
	t.Logf("B.Lat() = %v\n", b.Lat())
	t.Logf("B.Location = %v\n", b.Location)
}

// teste die DistanceTable
func TestDistanceTable(t *testing.T) {
	var from, to int
	for from = 0; from < len(GermanCities)-1; from++ {
		for to = from + 1; to < len(GermanCities); to++ {
			fmt.Printf("%v -> %v = %.0f km\n",
				GermanCities[from].Name(), GermanCities[to].Name(),
				CityDistanceTable[from<<8+to])
		}
	}
}

func TestCalculateRoundtrip(t *testing.T) {
	var end bool
	var min, max, dist float64
	min = 0.
	max = 0.
	start := time.Now()
	route := permute.NewPermutation(9)
	for {
		dist = CalculateRoundtrip(0, CityDistanceTable, route)
		if dist > max {
			max = dist
		}
		if min == 0. || dist < min {
			min = dist
		}
		//fmt.Printf("%v = %.0f km\n",
		//	ListRoundtrip(0, GermanCities, route),
		//	CalculateRoundtrip(0, CityDistanceTable, route))
		route, end = route.Next()
		if end {
			break
		}
	}
	done := time.Now()
	fmt.Printf("\nRoundtripCalculation took %v \n", done.Sub(start))
	fmt.Printf("Max: %.0f, Min: %.0f \n", max, min)
}
