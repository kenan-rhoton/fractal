package main

import (
	"github.com/gizak/termui"
	"os"
)

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	state := State{
		tasks:    NewTask("root"),
		selected: 0,
		list:     termui.NewList()}

	if len(os.Args) > 1 {
		state.Load(os.Args[1])
	}

	state.Init()

	//termui.UseTheme("helloworld")

	termui.Loop()

}
