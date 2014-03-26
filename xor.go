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
func commonPrefix(a, b *kadId) uint {
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
    return uint(i*8 + bit)
}

// Returns the ith bit of a kadId.
func (id *kadId) bit(i uint) uint8 {
    bit := i % 8
    byt := (i-bit)/8
    byte_id := id[byt]
    byte_id<<=bit
    byte_id>>=7
    return byte_id & 1
}

func (id *kadId) lessThan(greater *kadId) bool {
    for i := 0; i < 20; i++ {
        if greater[i] > id[i] {
            // We know greater is indeed greater since we start with
            // most significant bytes.
            return true
        } else if greater[i] != id[i] { // If both are equal, continue checking.
            return false
        }
    }
    return false // Ids are equal
}
