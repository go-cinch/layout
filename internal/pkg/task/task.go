package task

import (
	"context"

	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/worker"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

// ProviderSet is task providers.
var ProviderSet = wire.NewSet(New)

type Task struct {
	worker *worker.Worker
}

func (tk Task) Once(options ...func(*worker.RunOptions)) error {
	return tk.worker.Once(options...)
}

func (tk Task) Cron(options ...func(*worker.RunOptions)) error {
	return tk.worker.Cron(options...)
}

// New is initialize task worker from config
func New(c *conf.Bootstrap) (tk *Task, err error) {
	w := worker.New(
		worker.WithRedisURI(c.Data.Redis.Dsn),
		worker.WithGroup(c.Name),
		worker.WithHandler(func(ctx context.Context, p worker.Payload) error {
			return process(task{
				ctx:     ctx,
				payload: p,
			})
		}),
	)
	if w.Error != nil {
		log.Error(w.Error)
		err = errors.New("initialize worker failed")
		return
	}

	tk = &Task{
		worker: w,
	}

	for id, item := range c.Task {
		err = tk.worker.Cron(
			worker.WithRunUUID(id),
			worker.WithRunGroup(item.Name),
			worker.WithRunExpr(item.Expr),
			worker.WithRunTimeout(int(item.Timeout)),
			worker.WithRunMaxRetry(int(item.Retry)),
		)
		if err != nil {
			log.Error(err)
			err = errors.New("initialize worker failed")
			return
		}
	}

	log.Info("initialize worker success")
	return
}

type task struct {
	ctx     context.Context
	payload worker.Payload
}

func process(t task) (err error) {
	tr := otel.Tracer("task")
	_, span := tr.Start(t.ctx, "Task")
	defer span.End()
	switch t.payload.Group {
	case "task1":
		// t.game.Get(ctx, 1)
	case "task2":
		// t.game.Get(ctx, 2)
	}
	return
}
