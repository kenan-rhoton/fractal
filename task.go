package main

type Task struct {
	Title     string
	parent    *Task
	Children  []*Task
	Completed bool
}

func NewTask(name string) *Task {
	return &Task{
		Title:     name,
		Children:  make([]*Task, 0),
		Completed: false,
		parent:    nil}
}

func (t *Task) Rename(name string) {
	t.Title = name
}

func (t *Task) Name() string {
	return t.Title
}

func (t *Task) Add(name string) {
	child := NewTask(name)
	child.parent = t
	t.Children = append(t.Children, child)
}

func (t *Task) SubTasks() []*Task {
	return t.Children
}

func (t *Task) TaskList() []string {
	res := make([]string, len(t.Children))
	for i, s := range t.Children {
		res[i] = s.Name()
	}
	return res
}

func (t *Task) Remove(i int) {
	t.Children = append(t.Children[:i], t.Children[i+1:]...)
}

func (t *Task) Parent() *Task {
	return t.parent
}
