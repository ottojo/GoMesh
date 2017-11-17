package main

import (
	"fmt"
	"github.com/ottojo/tui-go"
)

var logBox *tui.Box
var sidebar *tui.Box
var input *tui.Entry
var inputBox *tui.Box
var currentNodeView *tui.Box
var ui tui.UI

func initUi() {
	nodeList := tui.NewList()

	sidebar = tui.NewVBox(nodeList)
	sidebar.SetBorder(true)

	logBox = tui.NewVBox()
	logBox.SetBorder(true)

	input = tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox = tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	currentNodeView = tui.NewVBox(tui.NewScrollArea(logBox), inputBox)
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
		selectedNode = list.Selected()
		updateLogView(list.Selected())
	})

	ui = tui.New(root)
	ui.SetKeybinding("q", func() { ui.Quit() })

	ui.Update(func() {
		updateLogView(selectedNode)
	})
	if err := ui.Run(); err != nil {
		panic(err)
	}

}

func updateLogView(i int) {
	for logBox.GetChildrenCount() != 0 {
		logBox.Remove(0)
	}
	logBox.Append(tui.NewSpacer())
	for _, line := range swarm[i].Log {
		logBox.Append(tui.NewHBox(
			tui.NewLabel(line),
		))
	}

}
