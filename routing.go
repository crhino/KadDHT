/*
    Christopher Piraino
    Kademlia Protocol: Routing Table
*/
package kademlia


import (

)

type kNode struct {
    prefix int // Number of bits in common with node ID.
    left *kNode
    right *kNode
    bucket []*Node // bucket is nil unless the kNode is a leaf node.
}

func newKNode(p int) *kNode {
    return &kNode{
                    prefix: p,
                    bucket: make([]*Node, 0, 20),
                 }
}

type kTree struct {
    k int
    tree *kNode
}

func newKTree(k int) (*kTree, error) {
    tree := &kTree{
                    k: 20,
                    tree: newKNode(0),
                 }
    return tree, nil
}

func (node *kNode) leaf() bool {
    return node.bucket == nil
}

// search is a recursive function that will find the kNode in which
// the p == t.prefix.
func (t *kNode) search(p int) (*kNode, error) {
    return nil, ErrorNotImplemented
}

func (t *kTree) add(node *Node) error {
    return ErrorNotImplemented
}

// find returns a ptr to the node with kadId == key, or an error if not found.
func (t *kTree) find(id, key *kadId) (*Node, error) {
    //prefix := commonPrefix(id, key)
    return nil, ErrorNotImplemented
}
