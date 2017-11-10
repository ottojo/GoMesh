package main

import (
	"math/rand"
	"time"
	"github.com/marcusolsson/tui-go"
	"fmt"
)

var debug bool = false
var message_delay_ms int = 0

func init() {
	rand.Seed(time.Now().UnixNano())
}

var swarm Swarm

func main() {
	//i, _ := strconv.ParseInt(os.Args[1], 10, 32)
	swarm = createSwarm(15, 50) //TODO while true if conn>nodes

	swarm.init()

	swarm[0].Route(Message{0, TYPE_BROADCAST, []byte("This is a broadcast."),
		swarm[0].Mac, swarm[len(swarm)-1].Mac,
		[]MAC{}, 5})

	swarm[0].Route(Message{1, TYPE_ROUTED_SHORTEST_PATH, []byte("This is a routed message."),
		swarm[0].Mac, swarm[len(swarm)-1].Mac,
		[]MAC{}, 10})

	nodeList := tui.NewList()

	sidebar := tui.NewVBox(nodeList)
	sidebar.SetBorder(true)

	log := tui.NewVBox()
	log.SetBorder(true)

	/*for _, m := range posts {
		log.Append(tui.NewHBox(
			tui.NewLabel(m.time),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.username))),
			tui.NewLabel(m.message),
			tui.NewSpacer(),
		))
	}*/

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	currentNodeView := tui.NewVBox(tui.NewScrollArea(log), inputBox)
	currentNodeView.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {

	})

	root := tui.NewHBox(sidebar, currentNodeView)

	var strings []string
	for _, node := range swarm {
		strings = append(strings, fmt.Sprintf("[%X] %s", int(node.Mac), node.Name))
	}

	nodeList.AddItems(strings...)
	nodeList.SetFocused(true)
	nodeList.SetSelected(0)
	nodeList.OnSelectionChanged(func(list *tui.List) {
		for log.GetChildCount() != 0 {
			log.Remove(0)
		}
		log.Append(tui.NewSpacer())
		for _, line := range swarm[list.Selected()].Log {
			log.Append(tui.NewHBox(
				tui.NewLabel(line),
			))
		}
	})

	ui := tui.New(root)
	ui.SetKeybinding("q", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}

}
