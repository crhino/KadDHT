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

func TestCommonPrefix(t *testing.T) {
    a := new(kadId)
    b := new(kadId)

    a[0] = 9
    b[0] = 8

    bit := commonPrefix(a, b)
    // a = 0b00001001, b = 0b00001000, thus commonPrefix should return bit 7.
    if bit != 7 {
        t.Errorf("The common prefix of a and b should be bit 7, but returned %v.", bit)
    }

    a[0] = 8
    b[0] = 8
    a[5] = 9
    b[5] = 8

    bit = commonPrefix(a, b)
    // a = 0b00001001, b = 0b00001000 of the 6th byte, 
    // thus commonPrefix should return bit == (5*8 + 7) == 47.
    if bit != 47 {
        t.Errorf("The common prefix of a and b should be bit 47, but returned %v.", bit)
    }
}

func TestBit(t *testing.T) {
    a := new(kadId)
    bit := a.bit(10)
    if bit != 0 {
        t.Errorf("The 10th bit of the kadId should be 0, but is %v", bit)
    }
    bit = a.bit(159)
    if bit != 0 {
        t.Errorf("The 159th bit of the kadId should be 0, but is %v", bit)
    }
    a[3] = 9
    bit = a.bit(28)
    if bit != 1 {
        t.Errorf("The 28th bit of the kadId should be 1, but is %v", bit)
    }
}
