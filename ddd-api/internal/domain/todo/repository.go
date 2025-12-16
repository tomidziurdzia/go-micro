package todo

import "context"

type Repository interface {
	Save(ctx context.Context, t *Todo) error
	GetByID(ctx context.Context, id ID) (*Todo, error)
	List(ctx context.Context) ([]*Todo, error)
	MarkCompleted(ctx context.Context, id ID) (*Todo, error)
}
