package task

import (
	"context"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/worker"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

// ProviderSet is task providers.
var ProviderSet = wire.NewSet(NewTask)

type Task struct {
	worker *worker.Worker
}

func (tk Task) Once(options ...func(*worker.RunOptions)) error {
	return tk.worker.Once(options...)
}

func (tk Task) Cron(options ...func(*worker.RunOptions)) error {
	return tk.worker.Cron(options...)
}

// NewTask is initialize task worker from config
func NewTask(c *conf.Bootstrap, game *biz.GameUseCase) (tk *Task, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()
	w := worker.New(
		worker.WithRedisUri(c.Data.Redis.Dsn),
		worker.WithGroup(c.Name),
		worker.WithHandler(func(ctx context.Context, p worker.Payload) error {
			return process(task{
				ctx:     ctx,
				payload: p,
				game:    game,
			})
		}),
	)
	if w.Error != nil {
		err = errors.WithMessage(w.Error, "initialize worker failed")
		return
	}

	tk = &Task{
		worker: w,
	}

	if len(c.Tasks) > 0 {
		for _, item := range c.Tasks {
			err = tk.worker.Cron(
				worker.WithRunGroup(item.Category),
				worker.WithRunUuid(item.Uuid),
				worker.WithRunExpr(item.Expr),
				worker.WithRunTimeout(int(item.Timeout)),
				worker.WithRunMaxRetry(int(item.Retry)),
			)
			if err != nil {
				err = errors.WithMessage(err, "initialize worker failed")
				return
			}
		}
	}

	log.Info("initialize worker success")
	return
}

type task struct {
	ctx     context.Context
	payload worker.Payload
	game    *biz.GameUseCase
}

func process(t task) (err error) {
	tr := otel.Tracer("task")
	ctx, span := tr.Start(t.ctx, "Task")
	defer span.End()
	switch t.payload.Group {
	case "task1":
		t.game.Get(ctx, 1)
	case "task2":
		t.game.Get(ctx, 2)
	}
	return
}
