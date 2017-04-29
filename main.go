package main

import (
	"math/rand"
	"time"
)

const DEBUG = true
const MESSAGE_DELAY_MS = 500

type Swarm []Node

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	swarm := createSwarm(10, 10)
	swarm[0].Route(Message{0, TYPE_BROADCAST, []byte("Hi"),
			       swarm[0].mac, swarm[len(swarm)-1].mac,
			       []int{}, 5})
}

func createSwarm(nodes, connections int) Swarm {
	var newSwarm Swarm
	for i := 0; i < nodes; i++ {
		newSwarm = append(newSwarm, Node{mac: i})
	}
	newSwarm = createConnections(newSwarm, connections)
	return newSwarm
}

func getRandomWithout(lowerBorderInclusive, upperBorderExclusive, not int) int {
	var random int = rand.Intn(upperBorderExclusive-lowerBorderInclusive) + lowerBorderInclusive
	for random == not {
		random = rand.Intn(upperBorderExclusive-lowerBorderInclusive) + lowerBorderInclusive
	}
	return random
}
