package main

type Task struct {
	name      string
	parent    *Task
	subTasks  []*Task
	completed bool
}

func NewTask(name string) *Task {
	return &Task{
		name:      name,
		subTasks:  make([]*Task, 0),
		completed: false,
		parent:    nil}
}

func (t *Task) Rename(name string) {
	t.name = name
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Add(name string) {
	child := NewTask(name)
	child.parent = t
	t.subTasks = append(t.subTasks, child)
}

func (t *Task) SubTasks() []*Task {
	return t.subTasks
}

func (t *Task) TaskList() []string {
	res := make([]string, len(t.subTasks))
	for i, s := range t.subTasks {
		res[i] = s.Name()
	}
	return res
}

func (t *Task) Remove(i int) {
	t.subTasks = append(t.subTasks[:i], t.subTasks[i+1:]...)
}

func (t *Task) Parent() *Task {
	return t.parent
}
