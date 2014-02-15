/*  Christopher Piraino

    Kademlia Protocl

    xor_test.go - Testing suite for the xor.go file.
*/
package kademlia

import (
    "testing"
)

func TestXor(t *testing.T) {
    a := new(kadId)
    b := new(kadId)

    a[0] = 9
    b[0] = 9

    c := xor(a, b)
    for i := 0; i < len(c); i++ {
        if c[i] != 0 {
            t.Errorf("a == b, thus a xor b should be 0, actual value is %v.", c)
        }
    }

    a[0] = 8
    c = xor(a, b)
    if c[0] != 1 {
        t.Errorf("8 XOR 9 == 1, but xor(a, b) returns %v.", c[0])
    }
}
