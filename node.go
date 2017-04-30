package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Node struct {
	Mac            MAC
	neighbours     []*Node
	RoutedMessages []Message
	Paths          map[MAC]Path
}

func (n Node) String() string {
	s := "MAC: " + fmt.Sprint(n.Mac) + "\n" +
		strconv.Itoa(len(n.neighbours)) + " Links:\n"
	for _, nb := range n.neighbours {
		s += fmt.Sprint(nb.Mac) + "\n"
	}
	s += "Routes:\n"
	for mac, path := range n.Paths {
		s += "To " + fmt.Sprint(mac) +
			" over " + fmt.Sprint(path.NextNodeMAC) +
			" in " + strconv.Itoa(path.Distance) + " Hops\n"
	}
	return s
}

func (n *Node) initRoutes() {
	paths := make(map[MAC]Path)
	for _, nb := range n.neighbours {
		paths[nb.Mac] = Path{NextNodeMAC: nb.Mac, Distance: 1}
	}
	n.Paths = paths
}

func (n *Node) Route(message Message) {
	time.Sleep(time.Duration(message_delay_ms) * time.Millisecond)
	if len(message.Hops) >= message.MaxHops || containsMessage(n.RoutedMessages, message) {
		return
		//Hop limit exceeded or already routed this message (useful for efficient broadcasting)
		//Drop message
	}
	n.RoutedMessages = append(n.RoutedMessages, message)
	message.Hops = append(message.Hops, n.Mac)
	switch message.MessageType {
	case TYPE_BROADCAST:
		n.routeBroadcast(message)
		break
	case TYPE_ROUTED_SHORTEST_PATH:
		n.routeShortestPath(message)
		break
	case TYPE_ROUTING_ANNOUNCEMENT:
		n.routingAnnouncement(message)
		break
	default:
		log.Panic("Node \"" + fmt.Sprint(n) + "\" recieved a message of unknown type!")
	}
}

func (n *Node) broadcastTo(message Message, nodes []*Node) {
	for _, node := range nodes {
		if !containsMAC(message.Hops, node.Mac) {
			node.Route(message) //TODO async?!
		}
	}
}

func containsMAC(macs []MAC, mac MAC) bool {
	for _, n := range macs {
		if n == mac {
			return true
		}
	}
	return false
}

func containsNode(nodes []*Node, node *Node) bool {
	for _, n := range nodes {
		if n.Mac == node.Mac {
			return true
		}
	}
	return false
}

func (n *Node) Connect(newNeighbour *Node) {
	n.neighbours = append(n.neighbours, newNeighbour)
}

func connectNodes(n1, n2 *Node) {
	n1.Connect(n2)
	n2.Connect(n1)
}

func removeNode(nodes []*Node, node *Node) []*Node {
	for k, n := range nodes {
		if node.Mac == n.Mac {
			return append(nodes[:k], nodes[k+1:]...)
		}
	}
	return nodes
}

func getNodeWithMac(nodes []*Node, mac MAC) (*Node, error) {
	for _, node := range nodes {
		if node.Mac == mac {
			return node, nil
		}
	}
	return nil, errors.New("Mac not found")
}
