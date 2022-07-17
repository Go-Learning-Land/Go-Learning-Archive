# Chapter 3

### Linear Data Structures
* Lists
* Sets
* Tuples
* Stacks

### LinkedList

```
LinkedList is a sequence of nodes that have properties and a reference to the next node in
the sequence. It is a linear data structure that is used to store data. The data structure
permits the addition and deletion of components from any node next to another node. They
are not stored contiguously in memory, which makes them different arrays.
```

```go
type Node struct{
    property int
    nextNode *Node
    }
```
The LinkedList class has headNode pointer as its property.
```go
// LinkedList class
type LinkedList struct {
 headNode *Node
}
```
[See LinkedList Source Code](./linked_list.go)

### [Doubly Linked List](./doubly_linked_list.go)

```go
// Node class
type Node struct {
 property int
 nextNode *Node
 previousNode *Node
}
```

### [Sets](./set.go)

