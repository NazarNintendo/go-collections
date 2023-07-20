# Go collections

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/github.com/NazarNintendo/go-collections)

This repository contains a collection of packages for the most common collections
in Go. It is intended to be used as a dependency for other projects.

## Clone the project

```
$ git clone https://github.com/NazarNintendo/go-collections
$ cd go-collections
```
## [base](base/)

```
$ cd base
```
A base package that implements a Collection type and several trivial interfaces. 
The Collection is embedded and used as a base for all other collections.  
It uses a generic type `T`, which is defined in the package that imports it.

There is an exported property `Elements` that is a map with keys of `T` and `type EmptyType struct{}` values. The reason I have used an empty struct for a type is because it uses zero bytes and takes up no memory.

The Collection provides implementations for the following interfaces: 

* `Slicer`
  * `Size() int` - returns the size of the collection
  * `Slice() []T` - returns a slice of the collection
* `Container`
  * `Contains(T) bool` - returns true if the collection contains the element 


## [set](set/)

```
$ cd set
```
[Set](https://en.wikipedia.org/wiki/Set_(abstract_data_type)) on Wikipedia.  

An empty set can be created by calling `NewSet()` or by calling `NewSetfromSlice()` and passing in a slice of `T` values.

Operations covered:

* Union `A ∪ B`
* Intersection `A ∩ B`
* Difference `A \ B`
* Symmetric Difference `A ∆ B`

## Contributing

If you would like to contribute, please open an issue or a pull request.