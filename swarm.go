package main

import (
	"log"
	"github.com/ottojo/go-randomdata"
	"math/big"
)

type Swarm []Node

func (s Swarm) get(i int) *Node {
	return &s[i]
}

func (s Swarm) init() {
	for i := 0; i < len(swarm); i++ {
		s[i].initRoutes()
	}
	for _, n := range s {
		n.broadcastRoutesTo(n.neighbours)
	}

}

func createSwarm(nodes, connections int64) Swarm {
	var newSwarm Swarm
	for i := int64(0); i < nodes; i++ {
		newSwarm = append(newSwarm, Node{Mac: MAC(i), Name: randomdata.SillyName()})
	}
	newSwarm = createConnections(newSwarm, connections)
	return newSwarm
}

func createConnections(swarm Swarm, connections int64) Swarm {
	maxLinks := new(big.Int)
	maxLinks.Binomial(int64(len(swarm)), 2)
	if maxLinks.Cmp(big.NewInt(int64(connections))) == -1 {
		connections = maxLinks.Int64()
	}
	for i := int64(0); i < connections; i++ {
		n1 := getRandomWithout(0, len(swarm), -1)
		n2 := getRandomWithout(0, len(swarm), n1)
		if connected(swarm[n1], swarm[n2]) {
			i--
			continue
		}
		connectNodes(&swarm[n1], &swarm[n2])
		if debug {
			log.Printf("Connected Nodes %v and %v\n", n1, n2)
		}
	}
	return swarm
}

func connected(n1, n2 Node) bool {
	return containsNode(n1.neighbours, &n2) && containsNode(n2.neighbours, &n1)
}
