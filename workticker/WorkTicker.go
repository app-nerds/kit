package workticker

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
)

type WorkTickerConfig[T any] struct {
	Logger             *logrus.Entry
	NumWorkers         int
	RateLimitPerSecond int
	TickFrequency      time.Duration
	WorkErrorChan      chan error
	WorkConfiguration  WorkConfiguration[T]
}

type Ticker[T any] interface {
	AddWorkConfiguration(config WorkConfiguration[T])
	Run(ctx context.Context)
}

type WorkTicker[T any] struct {
	logger            *logrus.Entry
	numWorkers        int
	tickFrequency     time.Duration
	workErrorChan     chan error
	workConfiguration WorkConfiguration[T]

	workChan chan WorkItem[T]
	limiter  ratelimit.Limiter
}

/*
NewWorkTicker creates a new WorkTicker instance of type T. Example:

		workTicker := workticker.NewWorkTicker(workticker.WorkTickerConfig{
		  Logger: logrus.New().WithField("app", "example"),
		  NumWorkers: 10,
		  RateLimitPerSecond: 10,
		  TickFrequency: 5 * time.Second,
		  WorkErrorChan: errorChan,
	    WorkConfiguration: workticker.WorkConfiguration[MyData]{
	      Handler: handlerFunc,
	      Retriever: retrieverFunc,
	    },
		})

		ctx, cancel := context.WithCancel(context.Background())
		go workTicker.Run(ctx)

		// Wait for app to close or something...
		cancel()
*/
func NewWorkTicker[T any](config WorkTickerConfig[T]) *WorkTicker[T] {
	return &WorkTicker[T]{
		logger:            config.Logger,
		numWorkers:        config.NumWorkers,
		tickFrequency:     config.TickFrequency,
		workErrorChan:     config.WorkErrorChan,
		workConfiguration: config.WorkConfiguration,

		workChan: make(chan WorkItem[T]),
		limiter:  ratelimit.New(config.RateLimitPerSecond),
	}
}

func (wp *WorkTicker[T]) Run(ctx context.Context) {
	wp.logger.Info("starting worker pool...")
	wg := sync.WaitGroup{}

	for i := 0; i < wp.numWorkers; i++ {
		wg.Add(1)

		go func(workerID int) {
			wp.logger.Infof("starting worker %d", workerID)
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					wp.logger.Infof("closing worker %d...", workerID)
					break

				case workItem := <-wp.workChan:
					wp.logger.WithFields(logrus.Fields{
						"workerID": workerID,
					}).Info("received work item. calling handler...")

					var err error

					if err = workItem.Handler(workerID, workItem.Data, wp.limiter); err != nil {
						if wp.workChan != nil {
							wp.workErrorChan <- err
						}
					}
				}
			}
		}(i)
	}

	go func() {
		ticker := time.NewTicker(wp.tickFrequency)

		for {
			select {
			case <-ticker.C:
				workItem, err := wp.workConfiguration.Retriever(wp.workConfiguration.Handler)

				if err != nil {
					wp.logger.WithError(err).Error("error retrieving work")
					continue
				}

				wp.workChan <- workItem

			case <-ctx.Done():
				wp.logger.Info("closing worker pool ticker...")
				break
			}
		}
	}()

	wg.Wait()
	close(wp.workChan)
}
