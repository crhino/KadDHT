/* 
    Christopher Piraino
    Kademlia Protocol

    routing_test.go - Test functions for routing.go
*/
package kademlia

import (
    "testing"
    "fmt"
)

func TestTreeInit(t *testing.T) {
    id := new(kadId)
    tree, err := newKTree(20, id)
    if err != nil {
        t.Errorf("err should be nil, but is %v. tree is %v.", err, tree)
    }
    if tree.k != 20 {
        t.Errorf("tree.k should be 20, but is %v.", tree.k)
    }
    if tree.id != id {
        t.Errorf("tree.id should be %v, but is %v.", id, tree.id)
    }
    if tree.tree == nil {
        t.Errorf("tree.tree should be initialized and not nil.")
    }
    if tree.tree.pLow != 0 {
        t.Errorf("The low prefix should be 0, but is %v.", tree.tree.pLow)
    }
    if tree.tree.pHigh != 159 {
        t.Errorf("The high prefix should be 159, but is %v.", tree.tree.pHigh)
    }
    if tree.tree.left != nil || tree.tree.right != nil {
        t.Errorf("The left and right subtrees should be nil, but left is %v and right is %v.",
                    tree.tree.left, tree.tree.right)
    }
    if len(tree.tree.bucket) != 0 {
        t.Errorf("Inital bucket size should be 0, is %v.", len(tree.tree.bucket))
    }
    if c := cap(tree.tree.bucket); c != 20 {
        t.Errorf("Initial bucket capacity should be 20, is %v.", c)
    }
}

func TestTreeAddAndFind(t *testing.T) {
    nid := new(kadId)
    nid[0] = 8
    node := &Node{id: *nid}
    ourId := new(kadId)
    ourId[5] = 9
    fmt.Printf("nid: %v and ourId: %v\n", nid, ourId)
    fmt.Printf("common prefix of ID's is %v.\n", commonPrefix(nid, ourId))
    tree, err := newKTree(20, ourId)
    if err != nil {
        t.Errorf("err is %v, and tree is %v.", err, tree)
    }
    err = tree.add(node)
    if err != nil {
        t.Errorf("err of tree.add(node) should be nil, but is %v", err)
    }
}
