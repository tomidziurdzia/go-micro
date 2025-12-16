package todo

import "time"

type Todo struct {
	id        ID
	title     Title
	completed bool
	createdAt time.Time
}

func New(title Title) (*Todo, error) {
	id, err := NewID()
	if err != nil {
		return nil, err
	}
	return &Todo{
		id:        id,
		title:     title,
		completed: false,
		createdAt: time.Now().UTC(),
	}, nil
}

func (t *Todo) ID() ID              { return t.id }
func (t *Todo) Title() Title        { return t.title }
func (t *Todo) Completed() bool     { return t.completed }
func (t *Todo) CreatedAt() time.Time { return t.createdAt }

func (t *Todo) MarkCompleted() {
	t.completed = true
}
