package main

import "github.com/gizak/termui"

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
	state.Init()

	//termui.UseTheme("helloworld")

	termui.Loop()

}
