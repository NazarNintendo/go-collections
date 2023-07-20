// Package set implements a set data structure and common set theory operations.
package set

import (
	"fmt"
	"github.com/NazarNintendo/go-collections/base"
)

// Set is an implementation of a set data structure.
// For usage examples see the tests.
type Set[T comparable] struct {
	base.Collection[T]
}

// NewSet returns a new empty set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		base.Collection[T]{
			Elements: map[T]base.EmptyType{},
		},
	}
}

// NewSetFromSlice returns a new set with the elements of the slice.
// The slice is not modified and the duplicates are removed.
func NewSetFromSlice[T comparable](slice []T) *Set[T] {
	set := &Set[T]{
		base.Collection[T]{
			Elements: make(map[T]base.EmptyType, len(slice)),
		}}
	for _, element := range slice {
		set.Add(element)
	}
	return set
}

// Add adds the element to the set.
func (set *Set[T]) Add(element T) {
	if !set.Contains(element) {
		set.Elements[element] = base.EmptyType{}
	}
}

// Remove removes the element from the set.
func (set *Set[T]) Remove(element T) {
	delete(set.Elements, element)
}

// Union returns a new set with elements in the set or the others. The sets are not modified.
func (set *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range set.Elements {
		result.Add(element)
	}
	for element := range other.Elements {
		result.Add(element)
	}
	return result
}

// Intersection returns a new set with elements in the set and the others. The sets are not modified.
func (set *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range set.Elements {
		if other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// Difference returns a new set with elements in the set that are not in the others. The sets are not modified.
func (set *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range set.Elements {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// SymmetricDifference returns a new set with elements in the set or the others but not in both. The sets are not modified.
func (set *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range set.Elements {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	for element := range other.Elements {
		if !set.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// String returns a string representation of the set.
func (set *Set[T]) String() string {
	return fmt.Sprintf("%v", set.Slice())
}
