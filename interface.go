package typez

import "context"

type RunnableCtx interface {
	Run(ctx context.Context)
}

type Runnable interface {
	Run()
}
