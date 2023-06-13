package workerpoolv2

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
)

type WorkerPoolConfig struct {
	Logger             *logrus.Entry
	NumWorkers         int
	RateLimitPerSecond int
	TickFrequency      time.Duration
}

type Pool interface {
}

type WorkerPool struct {
	logger        *logrus.Entry
	numWorkers    int
	tickFrequency time.Duration

	workChan    chan WorkItemCollection
	workConfigs []WorkConfiguration
	limiter     ratelimit.Limiter
}

func NewWorkerPool(config WorkerPoolConfig) *WorkerPool {
	return &WorkerPool{
		logger:        config.Logger,
		numWorkers:    config.NumWorkers,
		tickFrequency: config.TickFrequency,

		workChan:    make(chan WorkItemCollection),
		workConfigs: make([]WorkConfiguration, 0, 10),
		limiter:     ratelimit.New(config.RateLimitPerSecond),
	}
}

func (wp *WorkerPool) AddWorkConfiguration(config WorkConfiguration) {
	wp.workConfigs = append(wp.workConfigs, config)
}

func (wp *WorkerPool) Run(ctx context.Context) {
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

				case workItems := <-wp.workChan:
					wp.logger.WithFields(logrus.Fields{
						"workerID":     workerID,
						"numWorkItems": len(workItems),
					}).Info("received work items. calling handler...")

					if len(workItems) > 0 {
						err := workItems.Handle(workerID, wp.limiter, workItems[0].Handler)

						if err != nil {
							wp.logger.WithError(err).WithFields(logrus.Fields{
								"workerID":     workerID,
								"numWorkItems": len(workItems),
							}).Error("error handling work items")
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
				for _, workConfig := range wp.workConfigs {
					workItems, err := workConfig.Retriever.RetrieveWork(workConfig.Handler)

					if err != nil {
						wp.logger.WithError(err).Error("error retrieving work")
						continue
					}

					wp.workChan <- workItems
				}

			case <-ctx.Done():
				wp.logger.Info("closing worker pool ticker...")
				break
			}
		}
	}()

	wg.Wait()
	close(wp.workChan)
}
