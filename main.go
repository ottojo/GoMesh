package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var debug bool = true
var message_delay_ms int = 0

func init() {
	rand.Seed(time.Now().UnixNano())
}

var swarm Swarm

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/nodes", nodeIndex)
	router.HandleFunc("/nodes/{nodeMAC}", nodeOverview)
	router.HandleFunc("/nodes/{nodeMAC}/send", nodeSend)
	router.HandleFunc("/nodes/{nodeMAC}/messages", nodeMessages)

	go http.ListenAndServe(":8080", router)

	swarm = createSwarm(10, 15)

	swarm.init()

	go swarm[0].Route(Message{0, TYPE_BROADCAST, []byte("This is a broadcast."),
		swarm[0].Mac, swarm[len(swarm)-1].Mac,
		[]MAC{}, 5})

	swarm[0].Route(Message{1, TYPE_ROUTED_SHORTEST_PATH, []byte("This is a routed message."),
		swarm[0].Mac, swarm[len(swarm)-1].Mac,
		[]MAC{}, 10})

	select {}
}

func handler(w http.ResponseWriter, r *http.Request) {

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func nodeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Nodes:")
}

func nodeOverview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	macString := vars["nodeMAC"]
	mac, _ := strconv.Atoi(macString)
	nodeJSON, err := json.Marshal(swarm[mac])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(nodeJSON))
}

func nodeSend(w http.ResponseWriter, r *http.Request) {}

func nodeMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	macString := vars["nodeMAC"]
	mac, _ := strconv.Atoi(macString)
	messagesJSON, err := json.Marshal(swarm[mac].RoutedMessages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(messagesJSON))
}
