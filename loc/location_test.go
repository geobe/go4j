package loc

import (
	"testing"
)

// teste die Entfernungsberechnungen
func TestDist(t *testing.T) {
	// San Francisco
	sfo := Location{lat: 37.8, lon: 122.4}
	// St. Petersburg
	stp := Location{lon: -30.3, lat: 59.9}
	// Hamburg
	hh := Location{53.5, 10.}
	// Berlin
	b := Location{52.5, 13.3}
	t.Logf("\nDist(SFO-STP) = %.0f km", sfo.Dist(stp))
	t.Logf("\nDist(HH-B) = %.0f km", hh.Dist(b))
}
