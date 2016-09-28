package permute

import (
	"fmt"
	"sort"
)

// eine permutation wird als slice von n
// ganzen zahlen dargestellt, die jeweils als index
// verwendet werden können.
type Permutation []int

// Methode für das Sort interface
func (p Permutation) Len() int {
	return len(p)
}

// Methode für das Sort interface
func (p Permutation) Less(i, j int) bool {
	return p[i] < p[j]
}

// Methode für das Sort interface
func (p Permutation) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Permutation) workcopy() Permutation {
	n := len(p)
	var arr []int = make([]int, n)
	var i int
	for i = 0; i < n; i++ {
		arr[i] = p[i]
	}
	return arr[:]
}

// eine permutation wird als slice von n ganzen zahlen dargestellt, die jeweils als index
// verwendet werden können.
func NewPermutation(n int, first ...int) (p Permutation) {
	var arr []int = make([]int, n)
	var i, start int
	if len(first) != 0 {
		start = first[0]
	} else {
		start = 1
	}
	for i = 0; i < n; i++ {
		arr[i] = i + start
	}
	p = arr[:]
	return
}

// berechne die nächste permutation und stelle fest, ob alle permutationen durchgeführt wurden.
// gib neue Permutation als kopie zurück. Außerdem ein flag, ob ende erreicht ist.
func (permutation Permutation) Next() (seq Permutation, end bool) {
	end = done(permutation)
	if !end {
		var (
			tail  Permutation
			found bool
			p     int
		)
		for pos := len(permutation) - 2; pos > 0; pos-- {
			tail = permutation[pos:] // vom ende her auf die Permutation sehen
			if done(tail) {
				p, found = xpos(tail, permutation[pos-1])
				// gibt es noch einen größeren wert im tail?
				if found {
					// zwei werte austauschen
					permutation[pos-1], permutation[pos+p] =
						permutation[pos+p], permutation[pos-1]
					sort.Sort(tail)
					break
				}
			} else if len(tail) == 2 {
				tail[0], tail[1] = tail[1], tail[0]
				break
			} else {
				panic("program should never come here")
			}
		}
	}
	seq = permutation //.copy()
	return
}

// Tausche ersten Wert so aus, dass eine Teilpermutation entsteht
func (permutation Permutation) PermuteFirst(pos int) (seq Permutation) {
	if int(pos) > len(permutation) {
		panic(fmt.Sprintf("illegal argument %d: pos too large", pos))
	}
	seq = permutation.workcopy()
	swap := seq[pos]
	for i := pos; i > 0; i-- {
		seq[i] = seq[i-1]
	}
	seq[0] = swap
	return
}

// nächste Permutation nur im hinteren Teil der Permutation durchführen
func (permutation Permutation) NextTail() (seq Permutation, end bool) {
	seq = permutation //.copy()
	tail := seq[1:]
	tail, end = tail.Next()
	for i := int(1); int(i) < len(seq); i++ {
		seq[i] = tail[i-1]
	}
	return
}

// überprüfe, ob die werte in slice s strikt kleiner werdend sind. Wenn ja, ist das ende
// der permutationen erreicht.
func done(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] < s[i+1] {
			return false
		}
	}
	return true
}

// suche nach der position des kleinsten wertes in slice s, der
// größer ist als der wert v und gib den positionsindex zurück
func xpos(s []int, v int) (p int, found bool) {
	var nv = s[0]
	p = 0
	for i := range s {
		if s[i] > v && s[i] < nv {
			nv = s[i]
			p = i
		}
	}
	found = nv > v
	return
}
