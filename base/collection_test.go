package base

import (
	"fmt"
	"testing"
)

func TestNewCollection(t *testing.T) {
	collection := NewCollection[int]()
	if collection.Elements == nil {
		t.Error("NewCollection should return a collection with elements")
	}
}

func TestSize(t *testing.T) {
	collection := NewCollection[int]()
	if collection.Size() != 0 {
		t.Error("Size should return 0 for an empty collection")
	}
	collection.Elements[1] = EmptyType{}
	if collection.Size() != 1 {
		t.Error("Size should return the number of elements in the collection")
	}
}

func TestSlice(t *testing.T) {
	collection := NewCollection[int]()
	slice := collection.Slice()
	if len(slice) != 0 {
		t.Error("Slice should return an empty slice for an empty collection")
	}
	collection.Elements[1] = EmptyType{}
	slice = collection.Slice()
	if len(slice) != 1 || slice[0] != 1 {
		t.Error("Slice should return a slice with the elements of the collection")
	}
}

func TestContains(t *testing.T) {
	collection := NewCollection[int]()
	if collection.Contains(1) {
		t.Error("Contains should return false for an empty collection")
	}
	collection.Elements[1] = EmptyType{}
	if !collection.Contains(1) {
		t.Error("Contains should return true for an element in the collection")
	}
	if collection.Contains(2) {
		t.Error("Contains should return false for an element not in the collection")
	}
}

func BenchmarkNewCollection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCollection[int]()
	}
}

func BenchmarkSize(b *testing.B) {
	collection := NewCollection[int]()
	for i := 0; i < b.N; i++ {
		collection.Size()
	}
}

func BenchmarkSlice(b *testing.B) {
	collection := NewCollection[int]()
	for i := 0; i < b.N; i++ {
		collection.Slice()
	}
}

func BenchmarkContains(b *testing.B) {
	collection := NewCollection[int]()
	for i := 0; i < b.N; i++ {
		collection.Contains(1)
	}
}

func ExampleNewCollection() {
	fmt.Println(NewCollection[int]())
	// Output: &{map[]}
}

func ExampleCollection_Size() {
	collection := NewCollection[int]()
	fmt.Println(collection.Size())
	// Output: 0
}

func ExampleCollection_Slice() {
	collection := NewCollection[int]()
	fmt.Println(collection.Slice())
	// Output: []
}

func ExampleCollection_Contains() {
	collection := NewCollection[int]()
	fmt.Println(collection.Contains(1))
	// Output: false
}
