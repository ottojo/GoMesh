package main

import (
	"time"
	"math/rand"
	"os"
	"strconv"
)

var debug bool = false
var message_delay_ms int = 0
var selectedNode = 0

func init() {
	rand.Seed(time.Now().UnixNano())
}

var swarm Swarm
var receiver = 1

func tick() {
	swarm[0].Route(Message{time.Now().Unix(), TYPE_ROUTED_SHORTEST_PATH, []byte("This is a routed message."),
		swarm[0].Mac, swarm[receiver].Mac,
		[]MAC{}, 10})
	receiver += 1;
	receiver = receiver % 9
}

func main() {
	//TODO routing while in UI?
	//TODO Progress/Activity indicator
	//TODO Routing messages to logBox
	//TODO Statistics window
	i := int64(10)
	if len(os.Args) >= 2 {
		i, _ = strconv.ParseInt(os.Args[1], 10, 32)
	}
	swarm = createSwarm(i, 20) //TODO while true if conn>nodes

	swarm.init()

	swarm[0].Route(Message{0, TYPE_BROADCAST, []byte("This is a broadcast."),
		swarm[0].Mac, swarm[len(swarm)-1].Mac,
		[]MAC{}, 5})

	startTime := time.Now().UnixNano()
	for time.Now().UnixNano() < startTime+(3 * time.Second).Nanoseconds() {

	}

	swarm[0].Route(Message{1, TYPE_ROUTED_SHORTEST_PATH, []byte("This is a routed message."),
		swarm[0].Mac, swarm[len(swarm)-1].Mac,
		[]MAC{}, 10})

	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for range ticker.C {
			tick()
		}
	}()

	ticker2 := time.NewTicker(time.Millisecond * 100)
	go func() {
		for range ticker2.C {
			ui.Update(func() {
				updateLogView(selectedNode)
			})
		}
	}()

	initUi()

}
