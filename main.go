package main

import (
	"github.com/gizak/termui"
)

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

func fetchLetter(path string) string {
	base := "/sys/kbd/"
	if len(base) > len(path) || base != path[0:len(base)] {
		return ""
	}

	return path[len(base):]
}

func (s *State) GetInput(title string) string {
	termui.ResetHandlers()
	par := termui.NewPar("")
	par.BorderLabel = title
	par.Width = (termui.TermWidth() * 3) / 4
	par.Height = 3
	par.X = (termui.TermWidth() / 2) - (par.Width / 2)
	par.Y = (termui.TermHeight() / 2) - 2
	par.TextBgColor = termui.ColorWhite
	par.TextFgColor = termui.ColorBlack
	par.Text += ""
	termui.Render(par)

	signal := make(chan bool)

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/<escape>", func(termui.Event) {
		signal <- true
	})

	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		signal <- true
	})

	termui.Handle("/sys/kbd/<space>", func(termui.Event) {
		par.Text += " "
		termui.Render(par)
	})
	// C-8 is Backspace for some reason
	termui.Handle("/sys/kbd/C-8", func(termui.Event) {
		par.Text = par.Text[:len(par.Text)-1]
		termui.Render(par)
	})

	termui.Handle("/sys/kbd", func(ev termui.Event) {
		par.Text += fetchLetter(ev.Path)
		termui.Render(par)
	})

	<-signal
	s.ResetHandlers()
	return par.Text
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

func (s *State) ResetHandlers() {
	termui.ResetHandlers()
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/a", func(termui.Event) {
		go s.Add()
	})
	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		subTasks := s.tasks.SubTasks()
		if len(subTasks) > 0 {
			s.tasks = subTasks[s.selected]
			s.selected = 0
		}
		s.Render()
	})
	termui.Handle("/sys/kbd/<delete>", func(termui.Event) {
		subTasks := s.tasks.SubTasks()
		if len(subTasks) > 0 {
			s.tasks.Remove(s.selected)
			if s.selected > 0 {
				s.selected -= 1
			}
		}
		s.Render()
	})
	// C-8 is Backspace for some reason
	termui.Handle("/sys/kbd/C-8", func(termui.Event) {
		if parent := s.tasks.Parent(); parent != nil {
			s.tasks = parent
			s.selected = 0
		}
		s.Render()
	})
	termui.Handle("/sys/kbd/<up>", func(termui.Event) {
		if s.selected > 0 {
			s.selected -= 1
		}
		s.Render()
	})
	termui.Handle("/sys/kbd/<down>", func(termui.Event) {
		if s.selected < len(s.list.Items)-1 {
			s.selected += 1
		}
		s.Render()
	})
	termui.Handle("/sys/wnd/resize", func(termui.Event) {
		s.Render()
	})
}

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
