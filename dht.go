/*
    Christopher Piraino
    
    Kademlia: Peer-to-Peer Protocol

*/
package kademlia


import (
    "net"
    "errors"
    "strconv"
    "log"
    "unsafe"
)

type RPC int64
const (
    KADPING RPC = iota
    KADSTORE
    KADNODE
    KADVALUE
    KADBOOTSTRAP
    KADSTOP
)

var (
    ErrorNotImplemented = errors.New("function not implemented")
    ErrorNotFound       = errors.New("Could not find value")
)

type Config struct {
    k int // System-wide variable used to determine the maximum number of buckets
    alpha int // System-wide variable used for concurrency.
}

type Node struct {
    Addr *net.UDPAddr
    Id kadId
}

type DHT struct {
    config Config
    Node *Node
    routing *kTree
}

// Depending on the RPC, value or nodes field might be empty.
type kad_packet struct {
    rand_id kadId // RPC's must echo back this random ID to protect against forgery.
    send_id kadId // ID of the sending Node.
    rpc RPC
    key kadId
    value []byte
    nodes []Node
}

/* 
    Bootstrap_DHT(host, port, naddr) creates a new DHT node
    and adds it to the existing network of which host:port is
    a part of. If naddr is nil, the DHT node is in a network
    with only itself. "udp://host:port" will be the address of the
    created node.
*/
func Bootstrap_DHT(host string, port int, naddr *net.UDPAddr) (*DHT, error) {
    var udp_addr *net.UDPAddr
    var err error
    addr := host + ":"
    temp_byte := strconv.AppendInt([]byte(addr), int64(port), 10)
    addr = string(temp_byte)
    udp_addr, err = net.ResolveUDPAddr("udp", addr)
    if err != nil {
        return nil, err
    }
    id := randKadId()
    tree, err := newKTree(20, &id)
    if err != nil {
        log.Fatalln("Bootstrap_DHT err: ", err)
    }

    // Initialize the DHT structure.
    dht := &DHT {
        config: Config {
            k: 20,
            alpha: 3,
        },
        Node: &Node {
            Addr: udp_addr,
            Id: id,
        },
        routing: tree,
    }
    if naddr == nil {
        return dht, nil
    }

    // Must connect to network before checking id value
    new_node := false
    for !new_node {
        nodes, err := dht.Find_Node(dht.Node.Id)
        if err != nil {
            return nil, err
        }
        node_new := true
        for _, node := range nodes {
            if id == node.Id {
                node_new = false
            }
        }
        if node_new { // Better way to do this?
            new_node = true
        } else {
            dht.Node.Id = randKadId()
        }
    }
    return dht, nil
}

func (dht *DHT) Start() error {
    conn, err := net.ListenUDP("udp", dht.Node.Addr)
    if err != nil {
        log.Println("DHT Start() err: ", err)
        return err
    }
    stop_chan := make(<-chan RPC)
    go dht.RPC_response_loop(conn, stop_chan)
    <-stop_chan
    return nil
}

/*
            Kademlia RPC Commands

    Ping checks to see if a node is online. 

    Store tells a node to store the (key, value) pair.

    Find_Node tells the given node to return the k nearest nodes
    to key that it knows about.

    Find_Value is the same as Find_Node, but if any node is storing
    a value for the given key, it returns the value instead.


    Find_Node and Find_Value are both used in the lookup algorithm.
*/
func (dht *DHT) Ping(node *Node) error {
    return ErrorNotImplemented
}

func (dht *DHT) Store(node *Node, key kadId, value []byte) error  {
    return ErrorNotImplemented
}

func (dht *DHT) Find_Node(key kadId) ([]*Node, error) {
    return nil, ErrorNotImplemented
}

func (dht *DHT) Find_Value(key kadId) (*[]byte, []*Node, error) {
    return nil, nil, ErrorNotImplemented
}

func (dht *DHT) lookup_closest_nodes(key kadId) ([]*Node, error) {
    return nil, ErrorNotImplemented
}

// More channels should be parameters.
func (dht *DHT) RPC_response_loop(udp_conn *net.UDPConn, stop_chan <-chan RPC) {
    // Use net.ReadFromUDP function.
    var packet *kad_packet
    packet_size := int(unsafe.Sizeof(packet))
    buf := make([]byte, packet_size)
    for {
        n, raddr, err := udp_conn.ReadFromUDP(buf)
        if err != nil {
            log.Println("RPC_response_loop: ReadFromUDP err: ", err)
            continue
        }
        if n != packet_size {
            log.Println("RPC_response_loop: did not read correct number of bytes (",
                        packet_size, "), read ", n, " bytes instead.")
            continue
        }
        packet = (*kad_packet)(unsafe.Pointer(&buf[0]))
        switch packet.rpc {
            case KADPING:
                break
            case KADSTORE:
                break
            case KADNODE:
                break
            case KADVALUE:
                break
            case KADBOOTSTRAP:
                // Send boostrapping node all nodes closest to it's ID in order
                // to check for duplicates.
                node_ptrs := dht.routing.k_nearest_nodes(&packet.send_id)
                break
            case KADSTOP:
                break
        }
    }
}
