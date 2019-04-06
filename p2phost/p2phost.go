package p2phost

import (
	"bufio"
	"cli"
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"os"

	host "github.com/libp2p/go-libp2p-host"

	"github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	net "github.com/libp2p/go-libp2p-net"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
)

// Host - this is the variable to access the Host libp2p object anywhere
var Host host.Host

func handleStream(s net.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw)
	go writeData(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}
func readData(rw *bufio.ReadWriter) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
		}

	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		rw.WriteString(fmt.Sprintf("%s\n", sendData))
		rw.Flush()
	}

}

func InitP2P() {
	sourcePort := 8081
	dest := ""

	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		panic(err)
	}

	// 0.0.0.0 will listen on any interface device.
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", sourcePort))

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	Host, err = libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)

	if err != nil {
		panic(err)
	}

	cli.PostStatus("web", Host.ID().Pretty())

	if dest == "" {
		// Set a function as stream handler.
		// This function is called when a peer connects, and starts a stream with this protocol.
		// Only applies on the receiving side.
		Host.SetStreamHandler("/chat/1.0.0", handleStream)

		// Let's get the actual TCP port from our listen multiaddr, in case we're using 0 (default; random available port).
		var port string
		for _, la := range Host.Network().ListenAddresses() {
			if p, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
				port = p
				break
			}
		}

		if port == "" {
			panic("was not able to find actual local port")
		}

		// fmt.Printf("Run './chat -d /ip4/127.0.0.1/tcp/%v/p2p/%s' on another console.\n", port, host.ID().Pretty())
		// fmt.Println("You can replace 127.0.0.1 with public IP as well.")
		// fmt.Printf("\nWaiting for incoming connection\n\n")

		// Hang forever
		<-make(chan struct{})
	} else {
		fmt.Println("This node's multiaddresses:")
		for _, la := range Host.Addrs() {
			fmt.Printf(" - %v\n", la)
		}
		fmt.Println()

		// Turn the destination into a multiaddr.
		maddr, err := multiaddr.NewMultiaddr(dest)
		if err != nil {
			log.Fatalln(err)
		}

		// Extract the peer ID from the multiaddr.
		info, err := peerstore.InfoFromP2pAddr(maddr)
		if err != nil {
			log.Fatalln(err)
		}

		// Add the destination's peer multiaddress in the peerstore.
		// This will be used during connection and stream creation by libp2p.
		Host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

		// Start a stream with the destination.
		// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
		s, err := Host.NewStream(context.Background(), info.ID, "/chat/1.0.0")
		if err != nil {
			panic(err)
		}

		// Create a buffered stream so that read and writes are non blocking.
		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		// Create a thread to read and write data.
		go writeData(rw)
		go readData(rw)

		// Hang forever.
		// select {}
	}
}
