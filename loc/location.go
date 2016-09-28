package loc

import (
	"math"
)

// Geo-Punkt mit Breitengrad und Längengrad,
// wie aus GPS Daten oder Google Earth
type Location struct {
	lat float64
	lon float64
}

// Interface, das es ermöglicht, die Entfernung zwischen
// allen Objekten zu berechnen, die LatLon implementieren.
// Rückgabewerte:
// lat: Latitude
// lon: Longitude
type Locator interface {
	LatLon() (float64, float64)
}

// Umrechnung ° in radian
var toRad = math.Pi / 180

// Erdradius
const rEarth = 6371.

// Methode berechnet die Entfernung zwischen
// einer Location und einem Locator Objekt
func (l Location) Dist(lctr Locator) float64 {
	return dist(l, lctr)
}

// Function berechnet die Entfernung
// zwischen zwei Locator Objekten
func dist(lctr1, lctr2 Locator) float64 {
	lat1, lon1 := lctr1.LatLon()
	lat2, lon2 := lctr2.LatLon()
	phi1 := lat1 * toRad
	phi2 := lat2 * toRad
	dlambda := (lon1 - lon2) * toRad
	cosg := math.Sin(phi1)*math.Sin(phi2) +
		math.Cos(phi1)*math.Cos(phi2)*math.Cos(dlambda)
	return rEarth * math.Acos(cosg)
}

// implementiert interface Locator für den Typ Location
func (l Location) LatLon() (lat, lon float64) {
	lat = l.lat
	lon = l.lon
	return
}

// Getter Funktion für lat Wert
func (l *Location) Lat() float64 {
	return l.lat
}

// Setter Funktion für lat Wert
func (l *Location) SetLat(lat float64) {
	l.lat = lat
}

// Konstruktor
func New(lat, lon float64) Location {
	l := Location{lat, lon}
	return l
}
