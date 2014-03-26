/* 
    Christopher Piraino
    Kademlia Protocol

    routing_test.go - Test functions for routing.go
*/
package kademlia

import (
    "testing"
//    "fmt"
    "math/rand"
)

// The parent ID will be all 0's.
func NewTreeInit(t *testing.T) *kTree {
    id := new(kadId)
    tree, err := newKTree(20, id)
    if err != nil {
        t.Fatalf("err should be nil, but is %v. tree is %v.", err, tree)
    }
    if tree.k != 20 {
        t.Fatalf("tree.k should be 20, but is %v.", tree.k)
    }
    if tree.id != id {
        t.Fatalf("tree.id should be %v, but is %v.", id, tree.id)
    }
    if tree.tree == nil {
        t.Fatalf("tree.tree should be initialized and not nil.")
    }
    if tree.tree.pLow != 0 {
        t.Fatalf("The low prefix should be 0, but is %v.", tree.tree.pLow)
    }
    if tree.tree.pHigh != 159 {
        t.Fatalf("The high prefix should be 159, but is %v.", tree.tree.pHigh)
    }
    if tree.tree.left != nil || tree.tree.right != nil {
        t.Fatalf("The left and right subtrees should be nil, but left is %v and right is %v.",
                    tree.tree.left, tree.tree.right)
    }
    if len(tree.tree.bucket) != 0 {
        t.Fatalf("Inital bucket size should be 0, is %v.", len(tree.tree.bucket))
    }
    if c := cap(tree.tree.bucket); c != 20 {
        t.Fatalf("Initial bucket capacity should be 20, is %v.", c)
    }
    return tree
}

func AddNodeTest(t *testing.T, node *Node, tree *kTree) {
    err := tree.add(node)
    if err != nil {
        t.Errorf("err of tree.add(node) should be nil, but is %v", err)
    }
}

func TestTreeInit(t *testing.T) {
    tree := NewTreeInit(t)
    if tree == nil {
        t.Fatalf("This should not be allowed.")
    }
}

func TestTreeAddAndFind(t *testing.T) {
    nid := new(kadId)
    nid[0] = 8
    node := &Node{id: *nid}
    ourId := new(kadId)
    ourId[5] = 9
    tree := NewTreeInit(t)
    AddNodeTest(t, node, tree)
    findNode, err := tree.find(&node.id)
    if err != nil {
        t.Errorf("tree.find(node.id) returned an error: %v.", err)
    }
    if node != findNode {
        t.Errorf("node: %v != findNode: %v.", node, findNode)
    }
}

func TestSplitAddAndFind(t *testing.T) {
    tree := NewTreeInit(t)
    nodes := make([]*Node, 30, 30)
    for i := range nodes {
        nid := new(kadId)
        nid[(i+rand.Int()) % 20] = byte(i)
        nodes[i] = &Node{id: *nid}
        AddNodeTest(t, nodes[i], tree)
    }
    // Tree has parent ID of 0000...00.
    if tree.tree.left.pLow != 0 && tree.tree.left.pHigh != 0 {
        t.Errorf("The left subtree should hold only nodes with no common prefix.")
    }
    if tree.tree.right.pLow != 1 && tree.tree.right.pHigh != 159 {
        t.Errorf("The right subtree should hold only nodes with prefix btwn 1 and 159.")
    }
}

func nodeInSlice(n *Node, list []*Node) bool {
    for _, b := range list {
        if n == b {
            return true
        }
    }
    return false
}

func TestKNearestNodes(t *testing.T) {
    tree := NewTreeInit(t)
    nodes := make([]*Node, 30, 30)
    for i := range nodes {
        nid := new(kadId)
        nid[0] = byte(30-i)
        if i == 0 {
            nid[0] = 255 // prefix == 0
        }
        nodes[i] = &Node{id: *nid}
        AddNodeTest(t, nodes[i], tree)
    }
    nearest := tree.k_nearest_nodes(&nodes[0].id)
    if nearest == nil {
        t.Error("k_nearest_nodes should not return nil")
    }
    if len(nearest) != 20 {
        t.Errorf("Length of nearest should be 20, not %v.", len(nearest))
    }
    for i := range nearest {
        if !nodeInSlice(nodes[i], nearest) {
            t.Errorf("Node %v is not in the list. Id: %v", i, nodes[i].id)
        }
    }
}
