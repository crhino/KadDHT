/*  Christopher Piraino

    Kademlia Protocol

    xor.go - Implements the XOR distance operations necessary
    for the Kademlia protocol, for 160-bit IDs.
*/
package kademlia

import (
)

type kadId [20]byte

func xor(a, b *kadId) (c *kadId) {
    c = new(kadId)
    for i := 0; i < 20; i++ {
        c[i] = a[i] ^ b[i]
    }
    return
}


// Finds the first bit that differs between a and b,
// and returns that bit number.
func commonPrefix(a, b *kadId) int {
    i := 0
    for ; i < 20; i++ {
        if a[i] != b[i] {
            break
        }
    }

    if i == 20 {
        return 20*8 // a == b
    }

    xor := a[i] ^ b[i]
    bit := 0
    // 0b10000000 == 128
    for xor < 128 {
        xor<<=1
        bit++
    }
    return i*8 + bit
}
