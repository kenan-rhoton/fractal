package main

import "github.com/gizak/termui"

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
	termui.Handle("/sys/kbd/s", func(termui.Event) {
		s.Save()
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
	termui.Handle("/sys/kbd/<space>", func(termui.Event) {
		subTasks := s.tasks.SubTasks()
		if len(subTasks) > 0 {
			subTasks[s.selected].Completed = !subTasks[s.selected].Completed
		}
		s.Render()
	})
	termui.Handle("/sys/wnd/resize", func(termui.Event) {
		s.Render()
	})
}
