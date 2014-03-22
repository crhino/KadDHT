/*
    Christopher Piraino
    Kademlia Protocol: Routing Table
*/
package kademlia


import (
    "bytes"
)

// A Node in the routing table.
type kNode struct {
    // A kNode owns the ID range of ids with common pLow bits to common pHigh bits
    pLow uint // Low number of bits in common with node ID.
    pHigh uint // High number of bits in common with node ID.
    left *kNode
    right *kNode
    bucket []*Node // bucket is nil unless the kNode is a leaf node.
    overflow []*Node // Overflow bucket to hold nodes that should be refreshed.
}

func newKNode(pL, pH uint) *kNode {
    return &kNode{
                    pLow: pL,
                    pHigh: pH,
                    bucket: make([]*Node, 0, 20),
                    overflow: make([]*Node, 0, 0),
                 }
}

// A Tree data structure that implements the routing table.
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

// A Node is a leaf iff it's bucket is allocated to hold nodes.
func (node *kNode) leaf() bool {
    return node.bucket != nil
}

func (node *kNode) belongs(prefix uint) bool {
    return node.pLow <= prefix && node.pHigh >= prefix
}

// Adds a Node to the kNode. If len(n.bucket) == k, split the
// bucket.
func (n *kNode) add(k int, own_id *kadId, node *Node) {
    if len(n.bucket) < k {
        n.bucket = append(n.bucket, node)
        return
    }
    // If the bucket cannot be split, add node to overflow.
    if n.pLow == n.pHigh {
        n.overflow = append(n.overflow, node)
        return
    }
    // TODO: Push node down into subtrees if bucket is full
    pL_node := newKNode(n.pLow, n.pLow)
    pH_node := newKNode(n.pLow+1, n.pHigh)
    for _, n := range n.bucket {
        p := commonPrefix(own_id, &n.id)
        if pL_node.belongs(p) {
            pL_node.add(k, own_id, n)
        } else {
            pH_node.add(k, own_id, n)
        }
    }
    n.bucket = nil
    // Figure out which side of tree each node goes on.
    if own_id.bit(n.pLow) == 1 {
        n.left = pH_node
        n.right = pL_node
    } else {
        n.left = pL_node
        n.right = pH_node
    }
}

// search is a recursive function that will find the kNode with the
// closest prefix to p.
func (t *kNode) search(p uint) (*kNode, error) {
    if t.belongs(p) && t.leaf() {
        return t, nil
    }
    if t.right.belongs(p) {
        return t.right.search(p)
    } else if t.left.belongs(p) {
        return t.left.search(p)
    } else {
        return nil, ErrorNotFound
    }
}

// Adds the given node to the routing table.
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
    leaf.add(t.k, t.id, node)
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

func (t *kTree) k_nearest_nodes(key *kadId) ([]*Node) {
    return nil
}
