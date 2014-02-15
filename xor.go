/*  Christopher Piraino

    Kademlia Protocol

    xor.go - Implements the XOR distance operations necessary
    for the Kademlia protocol, for 160-bit IDs.
*/
package kademlia

import (
    //"bytes"
)

type kadId [20]byte

func xor(a, b *kadId) (c *kadId) {
    c = new(kadId)
    for i := 0; i < 20; i++ {
        c[i] = a[i] ^ b[i]
    }
    return
}


