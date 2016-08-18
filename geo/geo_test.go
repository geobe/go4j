package geo

import (
	"testing"
	"fmt"
	"github.com/geobe/go4j/permute"
	"time"
)

// teste die Entfernungsberechnungen
func TestDist(t *testing.T) {
	sfo := Location{lat: 37.8, lon: 122.4}
	stp := Location{lat: 59.9, lon: -30.3}
	hh := Location{lat: 53.5, lon: 10.}
	b := Location{lat: 52.5, lon: 13.3}
	t.Logf("\nDist(SFO-STP) = %.0f km", sfo.Dist(stp))
	t.Logf("\nDist(HH-B) = %.0f km", hh.Dist(b))
}

// teste, ob Dist auch für City Objekte funktioniert
func TestCities(t *testing.T) {
	for _, city := range GermanCities {
		fmt.Printf("Luftlinie von Zwickau nach %s sind %.0f km\n", city.name, Zwickau.Dist(city))
	}
}

// teste die DistanceTable
func xTestDistanceTable(t *testing.T) {
	var from, to int
	for from = 0; from < len(GermanCities) - 1; from++ {
		for to = from + 1; to < len(GermanCities); to++ {
			fmt.Printf("%v -> %v = %.0f km\n",
				GermanCities[from].name, GermanCities[to].name,
				CityDistanceTable[from << 8 + to])
		}
	}
}

func TestCalculateRoundtrip(t *testing.T) {
	var end bool
	var min, max, dist float64
	min = 0.
	max = 0.
	start := time.Now()
	route := permute.NewPermutation(11)
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
