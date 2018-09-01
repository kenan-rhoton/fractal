package main

import "github.com/gizak/termui"

type State struct {
	tasks    *Task
	list     *termui.List
	selected int
}

func (s *State) Init() {
	s.list.ItemFgColor = termui.ColorYellow
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(12, 0, s.list)))
	s.Render()
	s.ResetHandlers()
}

func (s *State) Add() {
	name := s.GetInput("Task Name")
	s.tasks.Add(name)
	s.Render()
}

func (s *State) Render() {
	s.list.Items = s.tasks.TaskList()
	if s.selected < len(s.list.Items) {
		s.list.Items[s.selected] = "[" + s.list.Items[s.selected] + "](bg-red)"
	}
	s.list.BorderLabel = s.tasks.Name()
	s.list.Height = termui.TermHeight()
	s.list.Width = termui.TermWidth()
	termui.Body.Align()
	termui.Render(s.list)
}