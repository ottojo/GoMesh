package main

import (
	"strconv"
	"fmt"
	"log"
)

const (
	TYPE_BROADCAST            = 0 //Message shall reach every single node in the network once
	TYPE_ROUTED_SHORTEST_PATH = 1 //Message shall reach it's destination with as few hops as possible
	TYPE_ROUTING_ANNOUNCEMENT = 2 //Node is broadcasting routing information
)

type Path struct {
	nextNode *Node
}

type Message struct {
	id          int
	messageType int
	payload     []byte
	senderMAC   int
	receiverMAC int
	hops        []int
	maxHops     int
}

func (m Message) String() string {
	s := "ID: " + strconv.Itoa(m.id) +
		" From: " + fmt.Sprint(m.senderMAC) +
		" To: " + fmt.Sprint(m.receiverMAC) +
		" over [" + sprintMACs(m.hops) +
		"] : " + string(m.payload)
	return s
}

func (m Message) Equals(m2 Message) bool {
	return m.id == m2.id && m.senderMAC == m2.senderMAC && m.receiverMAC == m2.receiverMAC
}

func sprintMACs(MACs []int) string {
	s := ""
	for _, MAC := range MACs {
		s += strconv.Itoa(MAC) + " "
	}
	return s
}

func printMessage(message Message) {
	log.Print(message)
}

func containsMessage(messages []Message, message Message) bool {
	for _, m2 := range messages {
		if message.Equals(m2) {
			return true
		}
	}
	return false
}
