package main

import "github.com/gizak/termui"

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
