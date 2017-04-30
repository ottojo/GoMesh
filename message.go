package main

import (
	"fmt"
	"strconv"
)

const (
	TYPE_BROADCAST            = 0 //Message shall reach every single node in the network once
	TYPE_ROUTED_SHORTEST_PATH = 1 //Message shall reach it's destination with as few Hops as possible
	TYPE_ROUTING_ANNOUNCEMENT = 2 //Node is broadcasting routing information
)

type Message struct {
	Id          int64
	MessageType int
	Payload     []byte
	SenderMAC   MAC
	ReceiverMAC MAC
	Hops        []MAC
	MaxHops     int
}

func (m Message) String() string {
	s := "ID: " + strconv.FormatInt(m.Id, 16) +
		" From: " + fmt.Sprint(m.SenderMAC) +
		" To: " + fmt.Sprint(m.ReceiverMAC) +
		" over [" + sprintMACs(m.Hops) +
		"] : " + string(m.Payload)
	return s
}

func (m Message) Equals(m2 Message) bool {
	return m.Id == m2.Id && m.SenderMAC == m2.SenderMAC && m.ReceiverMAC == m2.ReceiverMAC
}

func containsMessage(messages []Message, message Message) bool {
	for _, m2 := range messages {
		if message.Equals(m2) {
			return true
		}
	}
	return false
}
