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