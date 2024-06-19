	type Task struct {
		ID          string `datastore:"-" gae:"$id"`
		Description string
		Done        bool
	}

	type TaskRepo interface {
		Create(ctx context.Context, t *Task) (*Task, error)
	}

	type taskRepo struct {
		*datastore.Client
	}

	func (r *taskRepo) Create(ctx context.Context, t *Task) (*Task, error) {
		key := datastore.IncompleteKey("Task", nil)
		key, err := r.Client.Put(ctx, key, t)
		if err != nil {
			return nil, err
		}
		t.ID = key.Name
		return t, nil
	}  
