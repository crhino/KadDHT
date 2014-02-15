/*
    Christopher Piraino
    
    Kademlia: Peer-to-Peer Protocol

*/
package kademlia


import (
    "net"
    "errors"
)

var (
    ErrorNotImplemented = errors.New("function not implemented.")
)

type Config struct {
    Port int
    k int // System-wide variable used to determine the maximum number of buckets
    alpha int // System-wide variable used for concurrency.
}

type Node struct {
    addr net.UDPAddr
    id string
}

type DHT struct {
    config Config
    Node *Node
    conn *net.UDPConn

}

func (dht *DHT) Ping(node *Node) error {
    return ErrorNotImplemented
}

func (dht *DHT) Store(node *Node, key, value []byte) error  {
    return ErrorNotImplemented
}

func (dht *DHT) Find_Node(key []byte) ([]Node, error) {
    return nil, ErrorNotImplemented
}

func (dht *DHT) Find_Value(key []byte) ([]byte, []Node, error) {
    return nil, nil, ErrorNotImplemented
}
