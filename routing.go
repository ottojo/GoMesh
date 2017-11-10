package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

func (n *Node) routingAnnouncement(message Message) {
	var pathsAnnouncement map[MAC]Path
	err := json.Unmarshal(message.Payload, &pathsAnnouncement)
	(*n).Log = append((*n).Log, fmt.Sprintf("Received Routing Announcement: \"%v\"", pathsAnnouncement))
	if err == nil && pathsAnnouncement != nil {
		newPaths := n.Paths
		for mac, path := range pathsAnnouncement {
			//Update Paths
			if mac != n.Mac { //Path to self not needed
				if path.Distance < n.Paths[mac].Distance {
					(*n).Log = append((*n).Log, fmt.Sprintf("Shorter Path to %v over %v", mac, path.NextNodeMAC))
					newPaths[mac] = Path{NextNodeMAC: message.SenderMAC, Distance: path.Distance + 1}
				} else if _, ok := n.Paths[mac]; !ok {
					//Path is new
					(*n).Log = append((*n).Log, fmt.Sprintf("New Path to %v over %v", mac, path.NextNodeMAC))
					newPaths[mac] = Path{NextNodeMAC: message.SenderMAC, Distance: path.Distance + 1}

				} else {
					(*n).Log = append((*n).Log, fmt.Sprintf(
						"Path to %v over %v (%v) is not shorter than existing path over %v (%v)",
						mac, path.NextNodeMAC, path.Distance, n.Paths[mac].NextNodeMAC, n.Paths[mac].Distance))
				}
			}
		}

		if !reflect.DeepEqual(newPaths, n.Paths) {
			n.Paths = newPaths
			sender, err := getNodeWithMac(n.neighbours, message.SenderMAC)
			if err == nil {
				n.broadcastRoutesTo(removeNode(n.neighbours, sender)) //TODO async
			}
		}
	} else {
		log.Printf("[%v]Received broken Routing Announcement: %v", n.Mac, err)
	}

}

func (n *Node) broadcastRoutesTo(nodes []*Node) {
	payload, err := json.Marshal(n.Paths)
	if err != nil {
		log.Panic(err)
	}
	message := Message{Payload: payload,
		MessageType: TYPE_ROUTING_ANNOUNCEMENT,
		SenderMAC: n.Mac,
		ReceiverMAC: -1,
		Id: time.Now().Unix(),
		MaxHops: 1,
	}
	n.broadcastTo(message, nodes)
}

func (n *Node) routeBroadcast(message Message) {
	(*n).Log = append((*n).Log, fmt.Sprintf("Received Broadcast: \"%v\"", message))
	n.broadcastTo(message, n.neighbours)
}

func (n *Node) routeShortestPath(message Message) {
	if message.ReceiverMAC == n.Mac {
		(*n).Log = append((*n).Log, fmt.Sprintf("Received routed message: \"%v\"", fmt.Sprint(message)))
		return
	}
	path := n.Paths[message.ReceiverMAC]
	nextNode, err := getNodeWithMac(n.neighbours, path.NextNodeMAC)
	if err == nil {
		nextNode.Route(message)
	}
}
