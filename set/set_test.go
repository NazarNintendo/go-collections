package set

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	if s.Elements == nil {
		t.Error("NewSet should return a set with a non-nil elements map")
	}
}

func TestNewSetFromSlice(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1, 2, 2, 3})
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Error("NewSetFromSlice should return a set with the elements of the slice")
	}
}

func TestContains(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Error("Contains should return true if the set contains the element")
	}
	if s.Contains(4) {
		t.Error("Contains should return false if the set does not contain the element")
	}
}

func TestAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Error("Add should add the element to the set")
	}
	s.Add(1)
	if len(s.Elements) != 1 {
		t.Error("Add should not add the element to the set if it is already present")
	}
}

func TestRemove(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	s.Remove(1)
	if s.Contains(1) {
		t.Error("Remove should remove the element from the set")
	}
	s.Remove(1)
	if len(s.Elements) != 2 {
		t.Error("Remove should not remove any element from the set if it is not present")
	}
}

func TestUnion(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	o := NewSetFromSlice[int]([]int{2, 30, 40})
	u := s.Union(o)
	if len(u.Elements) != 5 {
		t.Error("Union should return a set with elements in the set or the others")
	}
	if !u.Contains(1) || !u.Contains(2) || !u.Contains(3) || !u.Contains(30) || !u.Contains(40) {
		t.Error("Union should return a set with elements in the set or the others")
	}
}

func TestIntersection(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	o := NewSetFromSlice[int]([]int{2, 30, 40})
	i := s.Intersection(o)
	if !i.Contains(2) {
		t.Error("Intersection should return a set with elements in the set and the others")
	}
	if i.Contains(1) || i.Contains(3) || i.Contains(30) || i.Contains(40) {
		t.Error("Intersection should not return a set with elements in the set or the others")
	}
}

func TestDifference(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	o := NewSetFromSlice[int]([]int{2, 30, 40})
	d := s.Difference(o)
	if !d.Contains(1) || !d.Contains(3) {
		t.Error("Difference should return a set with elements in the set that are not in the others")
	}
	if d.Contains(2) {
		t.Error("Difference should not return a set with elements in the set that are in the others")
	}
}

func TestSymmetricDifference(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	o := NewSetFromSlice[int]([]int{2, 30, 40})
	d := s.SymmetricDifference(o)
	if !d.Contains(1) || !d.Contains(3) || !d.Contains(30) || !d.Contains(40) {
		t.Error("SymmetricDifference should return a set with elements in the set or the others but not both")
	}
	if d.Contains(2) {
		t.Error("SymmetricDifference should not return a set with elements in the set and the others")
	}
}

func TestSet_Slice(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	if len(s.Slice()) != 3 {
		t.Error("Slice should return a slice with the elements of the set")
	}
}

func TestSet_String(t *testing.T) {
	s := NewSetFromSlice[int]([]int{1})
	if s.String() != "[1]" {
		t.Error("String should return a string representation of the set")
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

var slice = makeRange(0, 1000000)

func BenchmarkNewSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSet[int]()
	}
}

func BenchmarkNewSetFromSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSetFromSlice[int](slice)
	}
}

func BenchmarkContains(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	for i := 0; i < b.N; i++ {
		s.Contains(1)
	}
}

func BenchmarkAdd(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	for i := 0; i < b.N; i++ {
		s.Add(1)
	}
}

func BenchmarkRemove(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	for i := 0; i < b.N; i++ {
		s.Remove(1)
	}
}

func BenchmarkUnion(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	o := NewSetFromSlice[int](append(slice[1:], 1000001))
	for i := 0; i < b.N; i++ {
		s.Union(o)
	}
}

func BenchmarkIntersection(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	o := NewSetFromSlice[int](append(slice[1:], 1000001))
	for i := 0; i < b.N; i++ {
		s.Intersection(o)
	}
}

func BenchmarkDifference(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	o := NewSetFromSlice[int](append(slice[1:], 1000001))
	for i := 0; i < b.N; i++ {
		s.Difference(o)
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	o := NewSetFromSlice[int](append(slice[1:], 1000001))
	for i := 0; i < b.N; i++ {
		s.SymmetricDifference(o)
	}
}

func BenchmarkSlice(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	for i := 0; i < b.N; i++ {
		s.Slice()
	}
}

func BenchmarkString(b *testing.B) {
	s := NewSetFromSlice[int](slice)
	for i := 0; i < b.N; i++ {
		s.String()
	}
}

func ExampleNewSet() {
	fmt.Println(NewSet[int]())
	// Output: []
}

func ExampleNewSetFromSlice() {
	fmt.Println(NewSetFromSlice[int]([]int{2, 2}))
	// Output: [2]
}

func ExampleSet_Add() {
	s := NewSet[int]()
	s.Add(1)
	fmt.Println(s)
	// Output: [1]
}

func ExampleSet_Remove() {
	s := NewSetFromSlice[int]([]int{1, 2})
	s.Remove(1)
	fmt.Println(s)
	// Output: [2]
}

func ExampleSet_Union() {
	s := NewSetFromSlice[int]([]int{1})
	o := NewSetFromSlice[int]([]int{2})
	fmt.Println(s.Union(o))
	// Output: [1 2]
}

func ExampleSet_Intersection() {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	o := NewSetFromSlice[int]([]int{2, 30, 40})
	fmt.Println(s.Intersection(o))
	// Output: [2]
}

func ExampleSet_Difference() {
	s := NewSetFromSlice[int]([]int{1, 2, 3})
	o := NewSetFromSlice[int]([]int{2, 3, 40})
	fmt.Println(s.Difference(o))
	// Output: [1]
}

func ExampleSet_SymmetricDifference() {
	s := NewSetFromSlice[int]([]int{2, 30})
	o := NewSetFromSlice[int]([]int{2, 30, 1})
	fmt.Println(s.SymmetricDifference(o))
	// Output: [1]
}

func ExampleSet_String() {
	s := NewSetFromSlice[int]([]int{1})
	fmt.Println(s.String())
	// Output: [1]
}
