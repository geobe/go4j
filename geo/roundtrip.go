package geo

import (
	"github.com/geobe/go4j/permute"
	"bytes"
)

func CalculateRoundtrip(first int, dt map[int]float64, sequence permute.Permutation) (dist float64) {
	var previous int
	for i, val := range sequence {
		if i == 0 {
			dist += dt[first << 8 + int(val)]
		} else {
			dist += dt[previous << 8 + int(val)]
		}
		previous = int(val)
	}
	dist += dt[previous << 8 + first]
	return
}

func ListRoundtrip(first int, cities []City, sequence permute.Permutation) string {
	var buffer bytes.Buffer
	buffer.WriteString(cities[uint8(first)].name)
	buffer.WriteString(" -> ")
	for _, val := range sequence {
		buffer.WriteString(cities[val].name)
		buffer.WriteString(" -> ")
	}
	buffer.WriteString(cities[uint8(first)].name)
	return buffer.String()
}
