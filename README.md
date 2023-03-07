# Local Libp2p Network

This project was amid to provide a P2P connection between people in Iran restricted networks, but it can be used for any other purpose. The code has been forked from the original Hyprspace project which aims to provide the ability to create a network locally between nodes behind NATs and firewalls.
you can find the original project here: [Original Hyprspace Documentation](https://github.com/hyprspace/hyprspace#readme)


## Table of Contents
- [A Bit of Backstory](#a-bit-of-backstory)
- [Tests have been done](#tests-have-been-done)
- [Use Cases](#use-cases)
  - [File sharing](#file-sharing)
  - [Private local network](#private-local-network)
  - [Private chat](#private-chat)
  - [Private VPN](#private-vpn)
- [Tutorial](#tutorial)
  - [Establishing network with your own Bootstrap Node](#establishing-network-with-your-own-bootstrap-node)
    - [Download and Install GO](#download-and-install-go)
    - [Installing IPFS](#installing-ipfs)
    - [Initializing IPFS](#initializing-ipfs)
    - [Adding Bootstrap Node](#adding-bootstrap-node)
    - [Start The network](#start-the-network)
  - [Test your connection](#test-your-connection)
  - [Setting up a bootstrap node beyond a NAT](#setting-up-a-bootstrap-node-beyond-a-nat)


## A Bit of Backstory
In resent months the government in Iran tries to increase the restriction in Iran and stop the communication between people, they have blocked many social medias and messengers 
and are trying to block the internet completely, so this is our responsibility to try many ways to establish a secure network between people for people. 
We have found that a local network can be established inside the Iran which helps us to share files thanks to the Libp2p. So we have decided to create this project
to encourage developers to help us to create a secure network for people in Iran.

## Tests have been done

## Use Cases

### File sharing

### Private local network

### Private chat

### Private VPN

## Tutorial

### Establishing network with your own Bootstrap Node
you have to add your own custom bootstrap node to peers list in node.go file to use them instead of default bootstrap nodes, follow the steps below to add your bootstrap node instead of default bootstrap nodes

### Download and Install GO
follow the steps from this link [Download and install GO](https://go.dev/doc/install) to install the latest GO inside your system

### Installing IPFS
follow the steps from this link [Download and install IPFS official binary](https://docs.ipfs.tech/install/command-line/#install-official-binary-distributions) to get the ipfs command line tool on your system

### Initializing IPFS
you can initialize ipfs on you system by executing the following command:
```
ipfs init
```
this command will create the necessary files inside a directory named ```.ipfs``` inside of your home directory you can check the directory to see what is happening

### Adding Bootstrap Node
first you need to remove all default bootstrap nodes from ipfs by executing this command:
```
ipfs bootstrap rm --all 
```
now you need to find your peer ip by executing this command:
```
hostname -I
```
then you have to find your peer id with this command:
```
ipfs config show | grep "PeerID"
```
finally add the bootstrap node to ipfs with this command
```
ipfs bootstrap add /ip4/<ip address of bootnode>/tcp/4001/ipfs/<peer identity hash of bootnode>
```

### Start The network
you can start ipfs daemon in background with this command:
```
ipfs daemon & 
```

### Test your connection
if you did all steps right then you can see the other nodes ip and id which are connected to the network by executing this command:
```
ipfs swarm peers
```

### Setting up a bootstrap node beyond a NAT

### configure your router 
you have to do a port forwarding or setting up virtual server with your router by port forwarding the 4001 port to the ip of bootstrap node 

### configure ipfs 
open the config file inside .ipfs directory and update like this:
```
"AppendAnnounce": [
      "/ip4/<ip-of-bootstrap-node>/tcp/4001",
      "/ip4/<ip-of-bootstrap-node>/udp/4001/quic",
      "/ip4/<ip-of-bootstrap-node>/udp/4001/quic/webtransport"
    ],
```