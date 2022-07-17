# Chapter 2 

## Table of Contents
* Arrays
* Slices
* Two-dimensional slices
* Maps
* Database operations
* Variadic functions
* CRUD web forms 
#### Requirements for Working
```bash
go get -u github.com/go-sql-driver/mysql
```

### Arrays

* [Book Source Code](./arrays.go)

### Slice

* [Go Slice : Usage and Internals](https://go.dev/blog/slices-intro)
  
```
Slices are passed by referring to functions. Big slices can be passed to functions without
impacting performance. Passing a slice as a reference to a function is demonstrated in the
code as follows
```
* [Slice Source code](./slice.go)
* [2D Slice Source code](./twodslices.go)

#### [Golang Make function](https://linuxhint.com/golang-make-function/)

### Maps and Database Operations

* [Map Source Code](./maps.go)
* [Database Operations](./database_operations.go)

### CRUD Webforms

* [Webforms Source Code](./webforms.go)
* [Webforms Templates](./templates/)
* 