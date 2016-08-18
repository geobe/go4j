package permute

import (
	"testing"
	//"fmt"
	"fmt"
)

func TestDone(t *testing.T) {
	slice1 := []uint8{5, 4, 3, 2, 1, 0}
	slice2 := []uint8{5, 4, 2, 3, 1, 0}
	slice3 := []uint8{25, 14, 3, }

	if ! done(slice1) {
		t.Errorf("failed with slice1: done not recognized")
	}
	if done(slice2) {
		t.Errorf("failed with slice2: done falsely recognized")
	}
	if ! done(slice3) {
		t.Errorf("failed with slice3: done not recognized")
	}
}

func TestXpos(t *testing.T) {
	slice1 := []uint8{4, 6, 3, 1, 5}
	p, found := xpos(slice1, 2)
	if found && p != 2 {
		t.Errorf("position not found, was %d, value %d \n", p, slice1[p])
	}
	if ! found {
		t.Errorf("value not found\n")
	}
	p, found = xpos(slice1, 6)
	if found {
		t.Errorf("wrong position found, was %d, value %d \n", p, slice1[p])
	}
}

func TestPermute(t *testing.T) {
	fmt.Println("running TestPermute")
	var perm = NewPermutation(5)
	fmt.Println(perm)
	var count int
	for {
		_, done := perm.Next()
		//fmt.Println(seq)
		count++
		if done {
			break
		}
	}
	fmt.Println(perm)
	fmt.Println(count)
}

func TestPermuteFirst(t *testing.T) {
	fmt.Println("running TestPermuteFirst")
	perm := NewPermutation(5)
	fmt.Println(perm)
	for i := 1; i < len(perm); i++ {
		fmt.Println(perm.PermuteFirst(uint8(i)))
	}
}

func xTestPermuteTail(t *testing.T) {
	fmt.Println("running TestPermuteTail")
	perm := NewPermutation(5)
	for i := 1; i < len(perm); i++ {
		head := perm.PermuteFirst(uint8(i))
		fmt.Println(head)
inner:
		for {
			head, end := head.NextTail()
			if !end {
				fmt.Println(head)
			} else {
				break inner
			}
		}
	}
}