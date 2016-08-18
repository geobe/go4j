package geo

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
var rEarth = 6371.

// Methode berechnet die Entfernung zwischen einer Location und einem Locator Objekt
func (p Location) Dist(ptr2 Locator) float64 {
	return Dist(p, ptr2)
	//lat2, lon2 := ptr2.LatLon()
	//phi1 := p.lat * toRad
	//phi2 := lat2 * toRad
	//dlambda := (p.lon - lon2) * toRad
	//cosg := math.Sin(phi1)*math.Sin(phi2) + math.Cos(phi1)*math.Cos(phi2)*math.Cos(dlambda)
	//return rEarth * math.Acos(cosg)
}

// Function berechnet die Entfernung zwischen zwei Locator Objekten
func Dist(ptr1, ptr2 Locator) float64 {
	lat1, lon1 := ptr1.LatLon()
	lat2, lon2 := ptr2.LatLon()
	phi1 := lat1 * toRad
	phi2 := lat2 * toRad
	dlambda := (lon1 - lon2) * toRad
	cosg := math.Sin(phi1)*math.Sin(phi2) + math.Cos(phi1)*math.Cos(phi2)*math.Cos(dlambda)
	return rEarth * math.Acos(cosg)
}

// implementiert interface Locator für den Typ Location
func (p Location) LatLon() (lat, lon float64) {
	lat = p.lat
	lon = p.lon
	return
}
