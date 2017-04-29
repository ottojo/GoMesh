package main

import (
	"strconv"
	"log"
	"fmt"
	"time"
)

type Node struct {
	mac            int
	neighbours     []*Node
	routedMessages []Message
}

func (n *Node) String() string {
	return "MAC: " + strconv.Itoa(n.mac) + ", " + strconv.Itoa(len(n.neighbours)) + " links"
}

func (n *Node) Route(message Message) {
	time.Sleep(MESSAGE_DELAY_MS * time.Millisecond)
	//log.Printf("[%v]Received Message: \"%v\"", n.mac, message)
	if len(message.hops) >= message.maxHops || containsMessage(n.routedMessages, message) {
		return
		//Hop limit exceeded or already routed this message (useful for efficient broadcasting)
		//Drop message
	}
	n.routedMessages = append(n.routedMessages, message)
	message.hops = append(message.hops, n.mac)
	switch message.messageType {
	case TYPE_BROADCAST:
		n.broadcast(message)
		break
	case TYPE_ROUTED_SHORTEST_PATH:
		break
	case TYPE_ROUTING_ANNOUNCEMENT:
		break
	default:
		log.Panic("Node \"" + fmt.Sprint(n) + "\" recieved a message of unknown type!")
	}
}

func (n *Node) broadcast(message Message) {
	log.Printf("[%v]Received Broadcast: \"%v\"", n.mac, message)

	for _, neighbour := range n.neighbours {
		if !containsMAC(message.hops, neighbour.mac) {
			//log.Printf("[%v]Relaying Broadcast to \"%v\"", n.mac, neighbour)
			neighbour.Route(message) //TODO async?!
		}
	}
}

func (n *Node) routeShortestPath(message Message) {
	/*if message.receiverMAC == n.mac {
		printMessage(message)
		return
	}
	if len(message.hops) < message.maxHops {
		message.hops = append(message.hops, n)
		for _, neighbour := range n.neighbours {
			if !containsMAC(message.hops, neighbour.mac) {
				neighbour.broadcast(message) //TODO async?!
			}
		}
	}*/
}

func containsMAC(macs []int, mac int) bool {
	for _, n := range macs {
		if n == mac {
			return true
		}
	}
	return false
}

func createConnections(swarm Swarm, connections int) Swarm {
	for i := 0; i < connections; i++ {
		n1 := getRandomWithout(0, len(swarm), -1)
		n2 := getRandomWithout(0, len(swarm), n1)
		connectNodes(&swarm[n1], &swarm[n2])
		if DEBUG {
			log.Printf("Connected Nodes %v and %v\n", n1, n2)
		}
	}
	return swarm
}

func (n *Node) Connect(newNeighbour *Node) {
	n.neighbours = append(n.neighbours, newNeighbour)
}

func connectNodes(n1, n2 *Node) {
	n1.Connect(n2)
	n2.Connect(n1)
}
