package p2p

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-tcp-transport"
	ma "github.com/multiformats/go-multiaddr"
)

// Protocol is a descriptor for the Hyprspace P2P Protocol.
const Protocol = "/hyprspace/0.0.1"

// CreateNode creates an internal Libp2p nodes and returns it and it's DHT Discovery service.
func CreateNode(ctx context.Context, inputKey string, port int, handler network.StreamHandler) (node host.Host, dhtOut *dht.IpfsDHT, err error) {
	// Unmarshal Private Key
	privateKey, err := crypto.UnmarshalPrivateKey([]byte(inputKey))
	if err != nil {
		return
	}

	//ip6quic := fmt.Sprintf("/ip6/::/udp/%d/quic", port)
	ip4quic := fmt.Sprintf("/ip4/0.0.0.0/udp/%d/quic", port)

	//ip6tcp := fmt.Sprintf("/ip6/::/tcp/%d", port)
	ip4tcp := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)

	// Create libp2p node
	node, err = libp2p.New(
		libp2p.ListenAddrStrings(ip4quic, ip4tcp),
		libp2p.Identity(privateKey),
		libp2p.DefaultSecurity,
		libp2p.EnableHolePunching(),
		libp2p.EnableNATService(),
		libp2p.NATPortMap(),
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			idht, err := dht.New(ctx, h)
			return idht, err
		}),
		libp2p.Ping(true),
	)
	if err != nil {
		return
	}

	// Setup Hyprspace Stream Handler
	node.SetStreamHandler(Protocol, handler)

	// Create DHT Subsystem
	dhtOut = dht.NewDHTClient(ctx, node, datastore.NewMapDatastore())

	// Define Bootstrap Nodes.
	peers := []string{}

	// Convert Bootstap Nodes into usable addresses.
	BootstrapPeers := make(map[peer.ID]*peer.AddrInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return node, dhtOut, err
		}
		pii, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return node, dhtOut, err
		}
		pi, ok := BootstrapPeers[pii.ID]
		if !ok {
			pi = &peer.AddrInfo{ID: pii.ID}
			BootstrapPeers[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	var wg sync.WaitGroup
	lock := sync.Mutex{}
	count := 0
	wg.Add(len(BootstrapPeers))
	for _, peerInfo := range BootstrapPeers {
		go func(peerInfo *peer.AddrInfo) {
			defer wg.Done()
			err := node.Connect(ctx, *peerInfo)
			if err == nil {
				lock.Lock()
				count++
				lock.Unlock()

			}
		}(peerInfo)
	}
	wg.Wait()
	fmt.Println("[+] Number of connected nodes in version 2:", count)
	if count < 1 {
		return node, dhtOut, errors.New("unable to bootstrap libp2p node")
	}

	return node, dhtOut, nil
}
