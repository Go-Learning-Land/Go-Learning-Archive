# Chapter 4

## Non-Linear Data Structures

* Trees
* Tables
* Containers
* Hash functions

### Tree

```
space usage for a binary search tree is of the order O(n), whereas the insert, search,
and delete operations are of the order O(log n). A binary search tree consists of nodes with
properties or attributes:
```
* A key integer
* A value integer
* The leftNode and rightNode instances of TreeNode

```go
// TreeNode class
type TreeNode struct {
 key int
 value int
 leftNode *TreeNode
 rightNode *TreeNode
}
```

```go
// BinarySearchTree class
type BinarySearchTree struct {
 rootNode *TreeNode
 lock sync.RWMutex
}
```
[what is sync.RWMutex](https://medium.com/golangspec/sync-rwmutex-ca6c6c3208a0)
[Source Code Binary tree](./binary_search_tree.go)

### AVL Tree

* [Source Code AVL Tree](./avl_tree.go)

### Tables

```go
// Table Class
type Table struct {
 Rows []Row
 Name string
 ColumnNames []string
}

// Row Class
type Row struct {
 Columns []Column
 Id int
}
// Column Class
type Column struct {
 Id int
 Value string
 }

```
* [Table Source Code](./table.go)

### [Circular Linked List](./circular_list.go)

[Golang container/ring data type source code](https://cs.opensource.google/go/go/+/refs/tags/go1.18.4:src/container/ring/ring.go;drc=2580d0e08d5e9f979b943758d3c49877fb2324cb;l=14)
```

### Hash Functions