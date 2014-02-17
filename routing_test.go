/* 
    Christopher Piraino
    Kademlia Protocol

    routing_test.go - Test functions for routing.go
*/
package kademlia

import (
    "testing"
)

func TestTreeInit(t *testing.T) {
    tree, err := newKTree(20)
    if err != nil {
        t.Errorf("err should be nil, but is %v. tree is %v.", err, tree)
    }
    if tree.k != 20 {
        t.Errorf("tree.k should be 20, but is %v.", tree.k)
    }
    if tree.tree == nil {
        t.Errorf("tree.tree should be initialized and not nil.")
    }
    if tree.tree.prefix != 0 {
        t.Errorf("The intial prefix should be 0, but is %v.", tree.tree.prefix)
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
    tree, err := newKTree(20)
    if err != nil {
        t.Errorf("err is %v, and tree is %v.", err, tree)
    }
    nid := new(kadId)
    node := &Node{id: *nid}
    err = tree.add(node)
    if err != nil {
        t.Errorf("err of tree.add(node) should be nil, but is %v", err)
    }
}
