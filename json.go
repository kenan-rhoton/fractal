package main

import (
	"encoding/json"
	"io/ioutil"
)

func (t *Task) Adopt() {
	for _, c := range t.Children {
		c.parent = t
		c.Adopt()
	}
}

func (s *State) Save() {
	root := s.tasks
	for {
		if root.Parent() != nil {
			root = root.Parent()
		} else {
			break
		}
	}
	marshalled, _ := json.MarshalIndent(root, "", "    ")
	ioutil.WriteFile("tasks.json", marshalled, 0644)
}

func (s *State) Load(filename string) {
	root := Task{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &root)
	root.Adopt()
	s.tasks = &root
}
