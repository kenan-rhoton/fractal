package main

import "testing"

func TestTaskName(t *testing.T) {
	task := NewTask("Potato")
	if task.Name() != "Potato" {
		t.Errorf("%s is not \"Potato\"", task.Name())
	}
	task = NewTask("Tomato")
	if task.Name() != "Tomato" {
		t.Errorf("%s is not \"Tomato\"", task.Name())
	}
}

func TestTaskRename(t *testing.T) {
	task := NewTask("Potato")
	task.Rename("Tomato")
	if task.Name() != "Tomato" {
		t.Errorf("%s is not \"Tomato\"", task.Name())
	}
}

func TestSubTask(t *testing.T) {
	task := NewTask("Potato")
	task.Add("Tomato")
	if task.SubTasks()[0].Name() != "Tomato" {
		t.Errorf("%s is not \"Tomato\"", task.SubTasks()[0].Name())
	}
}

func TestTaskList(t *testing.T) {
	data := []string{"Tomato", "Veggie", "Alfalfa"}
	task := NewTask("Potato")
	for _, d := range data {
		task.Add(d)
	}
	for i, l := range task.TaskList() {
		if l != data[i] {
			t.Errorf("%s is not %s", l, data[i])
		}

	}
}

func TestTaskRemove(t *testing.T) {
	data := []string{"Tomato", "Veggie", "Alfalfa"}
	task := NewTask("Potato")
	for _, d := range data {
		task.Add(d)
	}
	task.Remove(0)
	task.Remove(1)
	if task.TaskList()[0] != "Veggie" {
		t.Errorf(
			"%v is not %v",
			task.TaskList(),
			[]string{"Veggie"})
	}
}

func TestTaskParent(t *testing.T) {
	task := NewTask("Potato")
	task.Add("Tomato")
	if task.SubTasks()[0].Parent().Name() != "Potato" {
		t.Errorf("%s is not \"Potato\"", task.SubTasks()[0].Parent().Name())
	}
}
