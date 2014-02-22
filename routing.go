/*
    Christopher Piraino
    Kademlia Protocol: Routing Table
*/
package kademlia


import (
    "bytes"
)

type kNode struct {
    // A kNode owns the ID range of ids with common pLow bits to common pHigh bits
    pLow int // Low number of bits in common with node ID.
    pHigh int // High number of bits in common with node ID.
    left *kNode
    right *kNode
    bucket []*Node // bucket is nil unless the kNode is a leaf node.
}

func newKNode(pL, pH int) *kNode {
    return &kNode{
                    pLow: pL,
                    pHigh: pH,
                    bucket: make([]*Node, 0, 20),
                 }
}

type kTree struct {
    k int
    id *kadId
    tree *kNode
}

func newKTree(_k int, _id *kadId) (*kTree, error) {
    tree := &kTree{
                    k: _k,
                    id: _id,
                    tree: newKNode(0, 159),
                 }
    return tree, nil
}

func (node *kNode) leaf() bool {
    return node.bucket != nil
}

func (node *kNode) belongs(prefix int) bool {
    return node.pLow <= prefix && node.pHigh >= prefix
}

// Adds a Node to the kNode. If len(n.bucket) == 20, the Node is
// pushed down into a subtree of the kNode.
func (n *kNode) add(node *Node) {
    n.bucket = append(n.bucket, node)
    // TODO: Push node down into subtrees if bucket is full
}

// search is a recursive function that will find the kNode with the
// closest prefix to p.
func (t *kNode) search(p int) (*kNode, error) {
    if t.belongs(p) {
        return t, nil
    }
    if t.leaf() {
        return nil, ErrorNotFound
    }


    pNode, err := t.right.search(p)
    if err != nil {
        pNode, err = t.left.search(p)
        if err != nil {
            return nil, err
        }
    }
    return pNode, nil
}

func (t *kTree) add(node *Node) error {
    p := commonPrefix(t.id, &(node.id))
    if p == 160 {
        // DHT tried to add itself to the routing table.
        return nil
    }
    leaf, err := t.tree.search(p)
    if err != nil {
        return err
    }
    leaf.add(node)
    return nil
}

// find returns a ptr to the node with kadId == key, or an error if not found.
func (t *kTree) find(key *kadId) (*Node, error) {
    prefix := commonPrefix(t.id, key)
    node, err := t.tree.search(prefix)
    if err != nil {
        return nil, err
    }
    for i := range node.bucket {
        if bytes.Equal(node.bucket[i].id[:], key[:]) { //Better way to do this?
            return node.bucket[i], nil
        }
    }
    return nil, ErrorNotFound
}
