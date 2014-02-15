/*
    Christopher Piraino
    Kademlia Protocol: Routing Table
*/
package kademlia


import (

)

type kNode struct {
    left *kNode
    right *kNode
    bucket *[]Node // bucket is nil unless the kNode is a leaf node.
}

type kTree struct {
    k int
    tree *kNode
}

func (node *kNode) leaf() bool {
    return node.left == nil && node.right == nil
}
