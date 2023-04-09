# Instructions

- [Instructions](#instructions)
  - [For `ADSL` networks](#for-adsl-networks)
    - [Setting up your Modem](#setting-up-your-modem)
    - [Download and Install `Go`](#download-and-install-go)
    - [Download and Install `IPFS`](#download-and-install-ipfs)
    - [Initializing IPFS Node](#initializing-ipfs-node)
    - [Setting up `Bootstrap` Node](#setting-up-bootstrap-node)
    - [Start the network](#start-the-network)
  - [Setup IPFS config file behind `NAT`](#setup-ipfs-config-file-behind-nat)
    - [Setting up](#setting-up)
  
----------

***`First of all, you should leased a Public IP from your ISP`***

----------

## For `ADSL` networks

### Setting up your Modem

Enable port forwarding for port 4001 (TCP/UDP) toward your local IPFS node and set firewall corresponding settings.

### Download and Install `Go`

1) sudo apt-get update
2) wget https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
3) sudo tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
4) export PATH=$PATH:/usr/local/go/bin
5) source $HOME/.profile

### Download and Install `IPFS`

1) sudo apt-get update
2) wget https://dist.ipfs.io/go-ipfs/v0.17.0/go-ipfs_v0.17.0_linux-amd64.tar.gz
3) tar xvfz go-ipfs_v0.17.0_linux-amd64.tar.gz
4) sudo mv go-ipfs/ipfs /usr/local/bin/ipfs
5) For verify:
   - ipfs version

### Initializing IPFS Node

- IPFS_PATH=~/.ipfs ipfs init

### Setting up `Bootstrap` Node

1) IPFS_PATH=~/.ipfs ipfs bootstrap rm --all
2) IPFS_PATH=~/.ipfs ipfs config show | grep "PeerID"
3) IPFS_PATH=~/.ipfs ipfs bootstrap add /ip4/<your public ip address>/tcp/4001/p2p/<peer identity hash of bootnode>
- Example:
IPFS_PATH=~/.ipfs ipfs bootstrap add /ip4/172.25.10.5/tcp/4001/p2p/QmdbaLZsKA94tsYeKJEPyLThWARFCtWyJWuudBUd4z9KBU

*Run steps 1 and 3 on other nodes.*

### Start the network

- IPFS_PATH=~/.ipfs ipfs daemon &
  
## Setup IPFS config file behind `NAT`

Change IPFS config file when your boot node is behind NAT.

### Setting up

1) Open IPFS config file with an editor:
   - nano .ipfs/config
2) Set your Public IP address in 'AppendAnnounce' section of 'config' file.

  "AppendAnnounce": [
      "/ip4/*`public ip`*/tcp/4001",
      "/ip4/*`public ip`*/udp/4001/quic",
      "/ip4/*`public ip`*/udp/4001/quic/webtransport"
    ]

- Restart IPFS node

- IPFS_PATH=~/.ipfs ipfs shutdown
- IPFS_PATH=~/.ipfs ipfs daemon &
