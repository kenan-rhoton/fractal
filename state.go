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

func (s *State) FormatList() {
	taskinfo := s.tasks.SubTasks()
	list := s.tasks.TaskList()
	for i := range list {
		list[i] += "- Started " + taskinfo[i].CreatedAt.Format("Jan 2 - 15:04")
		if taskinfo[i].Completed {
			list[i] += "- Completed " + taskinfo[i].CompletedAt.Format("Jan 2 - 15:04")
		}
		if i == s.selected {
			if taskinfo[i].Completed {
				list[i] = "[" + list[i] + "](bg-blue)"
			} else {
				list[i] = "[" + list[i] + "](bg-red)"
			}
		} else if s.tasks.SubTasks()[i].Completed {
			list[i] = "[" + list[i] + "](fg-blue)"
		}
	}
	s.list.Items = list
}

func (s *State) Render() {
	s.FormatList()
	s.list.BorderLabel = s.tasks.Name()
	s.list.Height = termui.TermHeight()
	s.list.Width = termui.TermWidth()
	termui.Body.Align()
	termui.Render(s.list)
}
